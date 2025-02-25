package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/generated/obbr/accounts/models"
)

type OBBRBankClient struct {
	baseURL string
	*http.Client
}

func NewOBBRBankClient(config Config) BankClient {
	c := OBBRBankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankURL.String()

	return &c
}

func (c *OBBRBankClient) GetInternalAccounts(_ context.Context, subject string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.ResponseAccountList{}
		err      error
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/internal/accounts?id=%s", c.baseURL, subject), http.NoBody); err != nil {
		return InternalAccounts{}, err
	}

	if response, err = c.Client.Do(request); err != nil {
		return InternalAccounts{}, err
	}
	defer response.Body.Close()

	if bytes, err = ioutil.ReadAll(response.Body); err != nil {
		return InternalAccounts{}, err
	}

	if response.StatusCode != 200 {
		return InternalAccounts{}, fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(bytes))
	}

	if err = json.Unmarshal(bytes, &resp); err != nil {
		return InternalAccounts{}, nil
	}

	return c.ToInternalAccounts(resp), nil
}

func (c *OBBRBankClient) ToInternalAccounts(data models.ResponseAccountList) InternalAccounts {
	accounts := make([]InternalAccount, len(data.Data))
	for i, account := range data.Data {
		accounts[i] = InternalAccount{
			ID:   *account.AccountID,
			Name: *account.Number,
		}
	}
	return InternalAccounts{Accounts: accounts}
}
