package bankly_test

import (
	"crypto/tls"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetClientIDMtls(t *testing.T) {
	godotenv.Load(".env.test")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	scope := bankly.GetScope().PixAccountRead

	certificate, err := tls.LoadX509KeyPair(dir+"/cert/client.crt", dir+"/cert/client.key")
	if err != nil {
		log.Fatalf("could not load certificate: %v", err)
	}
	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), scope, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
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

}
