package bankly_test

import (
	"crypto/tls"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetClientIDMtls(t *testing.T) {
	godotenv.Load(".env.test")

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd: %v", err)
		return
	}

	scope := bankly.GetScope().PixAccountRead

	certificate, err := tls.LoadX509KeyPair(dir+"/cert/client.crt", dir+"/cert/client.key")
	if err != nil {
		t.Errorf("could not load certificate: %v", err)
		return
	}
	client := bankly.NewClient(os.Getenv("ENV"))
	client.SetCertificateMtls(certificate)

	response, errApi, err := client.Mtls().RegisterClientID(&bankly.RequestRegisterClientID{
		GrantTypes:              []string{"client_credentials"},
		TlsClientAuthSubjectDn:  os.Getenv("BANKLY_SUBJECT_DN"),
		TokenEndpointAuthMethod: "tls_client_auth",
		ResponseTypes:           []string{"access_token"},
		CompanyKey:              os.Getenv("COMPANYKEY"),
		Scope:                   scope,
	})

	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}

	if response == nil {
		t.Error("response is null")
		return
	}
	client.SetCertificateMtls(certificate)

	responseToken, err := client.RequestToken(response.ClientID, "", scope, true)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}

	if responseToken == nil {
		t.Error("response is null")
		return
	}

}
