package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudentity/openbanking-quickstart/generated/obuk/accounts/models"
)

type OBUKBankClient struct {
	baseURL string
	*http.Client
}

func NewOBUKBankClient(config Config) BankClient {
	c := OBUKBankClient{}

	c.Client = &http.Client{}
	c.baseURL = config.BankURL.String()

	return &c
}

func (c *OBUKBankClient) GetInternalAccounts(_ context.Context, subject string) (InternalAccounts, error) {
	var (
		request  *http.Request
		response *http.Response
		bytes    []byte
		resp     = models.OBReadAccount6{}
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

func (c *OBUKBankClient) ToInternalAccounts(data models.OBReadAccount6) InternalAccounts {
	accounts := make([]InternalAccount, len(data.Data.Account))
	for i, account := range data.Data.Account {
		accounts[i] = InternalAccount{
			ID:   string(*account.AccountID),
			Name: string(account.Nickname),
		}
	}
	return InternalAccounts{Accounts: accounts}
}
