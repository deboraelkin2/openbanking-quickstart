package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	for k, v := range map[string]string{
		"ACP_URL":      "https://localhost:8443",
		"ACP_MTLS_URL": "https://acp:8443",
		"APP_HOST":     "localhost",
		"UI_URL":       "https://localhost:8091",
		"CERT_FILE":    "cert.pem",
		"KEY_FILE":     "key.pem",
		"TENANT":       "default",
		"SPEC":         "obbr",
		"BANK_URL":     "http://bank-br:8070",
	} {
		os.Setenv(k, v)
	}

	config, err := LoadConfig()
	require.NoError(t, err)

	require.Equal(t, 8091, config.Port)
	require.Equal(t, "/app/data/my.db", config.DBFile)
	require.Equal(t, "https://localhost:8443", config.ACPURL)
	require.Equal(t, "https://acp:8443", config.ACPInternalURL)
	require.Equal(t, "localhost", config.AppHost)
	require.Equal(t, "https://localhost:8091", config.UIURL)
	require.Equal(t, "cert.pem", config.CertFile)
	require.Equal(t, "key.pem", config.KeyFile)
	require.Equal(t, Spec("obbr"), config.Spec)
	require.Equal(t, "http://bank-br:8070", config.BankURL)
}
