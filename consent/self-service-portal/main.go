package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Spec string

const (
	OBUK Spec = "obuk"
	CDR  Spec = "cdr"
	OBBR Spec = "obbr"
	FDX  Spec = "fdx"
)

type Config struct {
	SystemClientID              string        `env:"SYSTEM_CLIENT_ID,required"`
	SystemClientSecret          string        `env:"SYSTEM_CLIENT_SECRET" envDefault:"PBV7q0akoP603rZbU0EFdxbhZ-djxF7FIVwyKaLnBYU"`
	SystemIssuerURL             *url.URL      `env:"SYSTEM_ISSUER_URL,required"`
	Timeout                     time.Duration `env:"TIMEOUT" envDefault:"5s"`
	OpenbankingWorkspaceID      string        `env:"OPENBANKING_SERVER_ID,required"`
	Spec                        Spec          `env:"SPEC,required"`
	BankURL                     *url.URL      `env:"BANK_URL,required"`
	RootCA                      string        `env:"ROOT_CA" envDefault:"/ca.pem"`
	CertFile                    string        `env:"CERT_FILE" envDefault:"/bank_cert.pem"`
	KeyFile                     string        `env:"KEY_FILE" envDefault:"/bank_key.pem"`
	Port                        int           `env:"PORT" envDefault:"8085"`
	LoginAuthorizationServerURL string        `env:"LOGIN_AUTHORIZATION_SERVER_URL,required"`
	LoginClientID               string        `env:"LOGIN_CLIENT_ID,required"`
	LoginAuthorizationServerID  string        `env:"LOGIN_AUTHORIZATION_SERVER_ID,required"`
	LoginTenantID               string        `env:"LOGIN_TENANT_ID,required"`
	IntrospectClientID          string        `env:"INTROSPECT_CLIENT_ID,required"`
	IntrospectClientSecret      string        `env:"INTROSPECT_CLIENT_SECRET" envDefault:"KThGH68f-gMC4cscGLFeOpIU4EYriYhKspOV9IwHbnw"`
	IntrospectIssuerURL         *url.URL      `env:"INTROSPECT_ISSUER_URL,required"`
	EnableTLSServer  			      bool 		      `env:"ENABLE_TLS_SERVER" envDefault:"true"`
	BankClientConfig            BankClientConfig
}

type BankClientConfig struct {
	URL          *url.URL `env:"BANK_URL,required"`
	AccountsURL  *url.URL `env:"BANK_ACCOUNTS_ENDPOINT"`
	TokenURL     string   `env:"BANK_CLIENT_TOKEN_URL"`
	ClientID     string   `env:"BANK_CLIENT_ID"`
	ClientSecret string   `env:"BANK_CLIENT_SECRET"`
	Scopes       []string `env:"BANK_CLIENT_SCOPES"`
	CertFile     string   `env:"BANK_CLIENT_CERT_FILE"`
	KeyFile      string   `env:"BANK_CLIENT_KEY_FILE"`
}

func (c *Config) SystemClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:     c.SystemClientID,
		ClientSecret: c.SystemClientSecret,
		IssuerURL:    c.SystemIssuerURL,
		Scopes:       []string{"manage_openbanking_consents", "view_clients"},
		Timeout:      c.Timeout,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}
}

func (c *Config) IntrospectClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:     c.IntrospectClientID,
		ClientSecret: c.IntrospectClientSecret,
		IssuerURL:    c.IntrospectIssuerURL,
		Scopes:       []string{"introspect_tokens"},
		Timeout:      c.Timeout,
		CertFile:     c.CertFile,
		KeyFile:      c.KeyFile,
		RootCA:       c.RootCA,
	}
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config           Config
	Client           acpclient.Client
	IntrospectClient acpclient.Client
	BankClient       BankClient
	ConsentClient    ConsentClient
}

func NewServer() (Server, error) {
	var (
		bankClient BankClient
		server     = Server{}
		err        error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = acpclient.New(server.Config.SystemClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}
	if server.IntrospectClient, err = acpclient.New(server.Config.IntrospectClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init introspect acp client")
	}

	switch server.Config.Spec {
	case OBUK:
		server.BankClient = NewOBUKBankClient(server.Config)
		server.ConsentClient = NewOBUKConsentImpl(&server)
	case OBBR:
		server.BankClient = NewOBBRBankClient(server.Config)
		server.ConsentClient = NewOBBRConsentImpl(&server)
	case CDR:
		if bankClient, err = NewCDRBankClient(server.Config); err != nil {
			return server, fmt.Errorf("failed to creating new CDR bank client %w", err)
		}
		server.BankClient = bankClient
		server.ConsentClient = NewCDRArrangementImpl(&server)
	case FDX:
		if bankClient, err = NewFDXBankClient(server.Config); err != nil {
			return server, fmt.Errorf("failed to creating new FDX bank client %w", err)
		}
		server.BankClient = bankClient
		server.ConsentClient = NewFDXConsentImpl(&server)
	default:
		return server, fmt.Errorf("unsupported spec %s", server.Config.Spec)
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())

	r.GET("/consents", s.ListConsents())
	r.DELETE("/consents/:id", s.RevokeConsent())

	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": s.Config.LoginAuthorizationServerURL,
			"clientId":               s.Config.LoginClientID,
			"authorizationServerId":  s.Config.LoginAuthorizationServerID,
			"tenantId":               s.Config.LoginTenantID,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("web/app/build/index.html")
	})

	if s.Config.EnableTLSServer {
		logrus.Debugf("running consent self service server tls")
		return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
	}

	logrus.Debugf("running consent self service server non-tls")
	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}

func main() {
	var (
		server Server
		err    error
	)

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
