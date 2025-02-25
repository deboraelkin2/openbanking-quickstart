package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	cdrBank "github.com/cloudentity/openbanking-quickstart/generated/cdr/client"
	fdxBank "github.com/cloudentity/openbanking-quickstart/generated/fdx/client"
	obbrAccounts "github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/client"
	obbrPayments "github.com/cloudentity/openbanking-quickstart/generated/obbr/payments/client"
	obukAccounts "github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/client"
	payments_client "github.com/cloudentity/openbanking-quickstart/generated/obuk/payments/client"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/oauth2/client/oauth2"
	oauth2Models "github.com/cloudentity/acp-client-go/clients/oauth2/models"
)

type Clients struct {
	AcpAccountsClient acpclient.Client
	AcpPaymentsClient acpclient.Client
	BankClient        BankClient
	ConsentClient     ConsentClient
}

type BankClient interface {
	GetAccounts(c *gin.Context, accessToken string, bank ConnectedBank) ([]Account, error)
	GetTransactions(c *gin.Context, accessToken string, bank ConnectedBank) ([]Transaction, error)
	GetBalances(c *gin.Context, accessToken string, bank ConnectedBank) ([]Balance, error)
	CreatePayment(c *gin.Context, data interface{}, accessToken string) (PaymentCreated, error)
}

type BankClientCreationFn func(Config) (BankClient, error)

type ConsentClient interface {
	CreateConsentExplicitly() bool
	CreateAccountConsent(c *gin.Context) (string, error)
	CreatePaymentConsent(c *gin.Context, req CreatePaymentRequest) (string, error)
	GetPaymentConsent(c *gin.Context, consentID string) (interface{}, error)

	UsePAR() bool
	DoPAR(c *gin.Context) (string, acpclient.CSRF, error)
	Signer
}

func (c *Clients) RenewAccountsToken(ctx context.Context, bank ConnectedBank) (*oauth2Models.TokenResponse, error) {
	var (
		resp      oauth2.TokenOK
		request   *http.Request
		response  *http.Response
		assertion string
		body      []byte
		err       error
	)

	values := url.Values{
		"client_id":     {c.AcpAccountsClient.Config.ClientID},
		"grant_type":    {"refresh_token"},
		"refresh_token": {bank.RefreshToken},
	}

	if c.AcpAccountsClient.Config.AuthMethod == acpclient.ClientSecretPostAuthnMethod && c.AcpAccountsClient.Config.ClientSecret != "" {
		values.Add("client_secret", c.AcpAccountsClient.Config.ClientSecret)
	}

	if c.AcpAccountsClient.Config.AuthMethod == acpclient.PrivateKeyJwtAuthnMethod {
		if assertion, err = c.AcpAccountsClient.GenerateClientAssertion(); err != nil {
			return nil, err
		}
		values.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		values.Add("client_assertion", assertion)
	}

	if request, err = http.NewRequest(http.MethodPost, c.AcpAccountsClient.Config.GetTokenURL(), strings.NewReader(values.Encode())); err != nil {
		return nil, errors.Wrapf(err, "failed to create token request")
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if response, err = c.AcpAccountsClient.DoRequest(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &resp.Payload); err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func InitClients(config Config,
	signerCreateFn SignerCreationFn,
	bankClientCreateFn BankClientCreationFn,
	consentClientCreateFn ConsentClientCreationFn,
) (Clients, error) {
	var (
		clients              = Clients{}
		acpAccountsWebClient acpclient.Client
		acpPaymentsWebClient acpclient.Client
		bankClient           BankClient
		signer               Signer
		consentClient        ConsentClient
		err                  error
	)

	if acpAccountsWebClient, err = NewAcpClient(config, "/api/callback"); err != nil {
		return clients, errors.Wrapf(err, "failed to create acp accounts client")
	}

	if acpPaymentsWebClient, err = NewAcpClient(config, "/api/domestic/callback"); err != nil {
		return clients, errors.Wrapf(err, "failed to create acp payments client")
	}

	if signerCreateFn != nil {
		if signer, err = signerCreateFn(config.KeyFile); err != nil {
			return clients, errors.Wrapf(err, "failed to create consent message signer for %s", config.Spec)
		}
	}

	if bankClient, err = bankClientCreateFn(config); err != nil {
		return clients, errors.Wrapf(err, "failed to create bank client for %s", config.Spec)
	}

	if consentClientCreateFn != nil {
		consentClient = consentClientCreateFn(acpAccountsWebClient, acpPaymentsWebClient, signer)
	}

	return Clients{
		AcpAccountsClient: acpAccountsWebClient,
		AcpPaymentsClient: acpPaymentsWebClient,
		BankClient:        bankClient,
		ConsentClient:     consentClient,
	}, nil
}

func NewAcpClient(cfg Config, redirect string) (acpclient.Client, error) {
	var (
		authorizeURL, issuerURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", cfg.ACPInternalURL, cfg.Tenant, cfg.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", cfg.ACPURL, cfg.Tenant, cfg.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s%s", cfg.UIURL, redirect)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	config := acpclient.Config{
		ClientID:                      cfg.ClientID,
		IssuerURL:                     issuerURL,
		AuthorizeURL:                  authorizeURL,
		RedirectURL:                   redirectURL,
		RequestObjectSigningKeyFile:   cfg.KeyFile,
		RequestObjectExpiration:       &requestObjectExpiration,
		Scopes:                        cfg.ClientScopes,
		Timeout:                       time.Second * 5,
		CertFile:                      cfg.CertFile,
		KeyFile:                       cfg.KeyFile,
		RootCA:                        cfg.RootCA,
		ClientAssertionSigningKeyFile: cfg.KeyFile,
	}

	if cfg.Spec == CDR {
		config.SkipClientCredentialsAuthn = true
		config.AuthMethod = acpclient.PrivateKeyJwtAuthnMethod
	}

	if cfg.Spec == FDX {
		config.SkipClientCredentialsAuthn = true
		config.AuthMethod = acpclient.TLSClientAuthnMethod
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

type OBUKClient struct {
	*obukAccounts.Accounts
	*payments_client.Payments
}

func NewOBUKClient(config Config) (BankClient, error) {
	var (
		c   = &OBUKClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		hc,
	)

	c.Accounts = obukAccounts.New(tr, nil)
	c.Payments = payments_client.New(tr, nil)

	return c, nil
}

type CDRClient struct {
	*cdrBank.Banking
}

func NewCDRClient(config Config) (BankClient, error) {
	var (
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return nil, err
	}

	tr := NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		http.DefaultClient,
	)
	return &CDRClient{
		cdrBank.New(tr, nil),
	}, nil
}

type OBBRClient struct {
	*obbrAccounts.Accounts
	*obbrPayments.Payments
}

func NewOBBRClient(config Config) (BankClient, error) {
	var (
		c   = &OBBRClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Accounts = obbrAccounts.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path+"/accounts/v1",
		[]string{u.Scheme},
		hc,
	), nil)
	c.Payments = obbrPayments.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path+"/payments/v1",
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}

type ConsentClientCreationFn func(acpclient.Client, acpclient.Client, Signer) ConsentClient

type OBUKConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}

func NewOBUKConsentClient(accountsClient, paymentsClient acpclient.Client, signer Signer) ConsentClient {
	return &OBUKConsentClient{accountsClient, paymentsClient, signer}
}

type OBBRConsentClient struct {
	Accounts acpclient.Client
	Payments acpclient.Client
	Signer
}

func NewOBBRConsentClient(accountsClient, paymentsClient acpclient.Client, signer Signer) ConsentClient {
	return &OBBRConsentClient{accountsClient, paymentsClient, signer}
}

type FDXBankClient struct {
	*fdxBank.Client
}

func NewFDXBankClient(config Config) (BankClient, error) {
	var (
		c   = &FDXBankClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.BankURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Client = fdxBank.New(NewHTTPRuntimeWithClient(
		u.Host,
		u.Path,
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}

type CDRConsentClient struct {
	ClientID     string
	ClientSecret string
	PublicClient acpclient.Client
}

func NewCDRConsentClient(publicClient, clientCredentialsClient acpclient.Client, _ Signer) ConsentClient {
	return &CDRConsentClient{
		ClientID:     clientCredentialsClient.Config.ClientID,
		ClientSecret: clientCredentialsClient.Config.ClientSecret,
		PublicClient: publicClient,
	}
}

func (c *CDRConsentClient) CreateConsentExplicitly() bool {
	return false
}

func (c *CDRConsentClient) UsePAR() bool {
	return true
}

func (c *CDRConsentClient) DoPAR(ctx *gin.Context) (string, acpclient.CSRF, error) {
	var (
		csrf acpclient.CSRF
		resp acpclient.PARResponse
		err  error
	)

	if resp, csrf, err = c.PublicClient.DoPAR(
		acpclient.WithResponseType("code id_token"),
		acpclient.WithPKCE(),
		acpclient.WithOpenbankingACR([]string{"urn:cds.au:cdr:3"}),
	); err != nil {
		return "", acpclient.CSRF{}, err
	}
	return resp.RequestURI, csrf, err
}

func (c *CDRConsentClient) CreateAccountConsent(ctx *gin.Context) (string, error) {
	return "", nil
}

func (c *CDRConsentClient) DoRequestObjectEncryption() bool {
	return false
}

func (c *CDRConsentClient) GetPaymentConsent(ctx *gin.Context, consentID string) (interface{}, error) {
	return "", nil
}

func (c *CDRConsentClient) CreatePaymentConsent(ctx *gin.Context, req CreatePaymentRequest) (string, error) {
	return "", nil
}

func (c *CDRConsentClient) Sign([]byte) (string, error) {
	return "", nil
}
