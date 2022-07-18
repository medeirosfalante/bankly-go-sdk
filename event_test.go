package bankly_test

import (
	"crypto/tls"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetEvent(t *testing.T) {
	godotenv.Load(".env.test")

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd: %v", err)
		return
	}

	certificate, err := tls.LoadX509KeyPair(dir+"/cert/client.crt", dir+"/cert/client.key")
	if err != nil {
		t.Errorf("could not load certificate: %v", err)
		return
	}
	client := bankly.NewClient(os.Getenv("ENV"))

	client.SetCertificateMtls(certificate)
	responseMtls, errApi, err := client.Mtls().RegisterClientID(&bankly.RequestRegisterClientID{
		GrantTypes:              []string{"client_credentials"},
		TlsClientAuthSubjectDn:  os.Getenv("BANKLY_SUBJECT_DN"),
		TokenEndpointAuthMethod: "tls_client_auth",
		ResponseTypes:           []string{"access_token"},
		CompanyKey:              os.Getenv("COMPANYKEY"),
		Scope:                   bankly.GetScope().EventsRead,
	})

	if err != nil {
		t.Errorf("err : responseMtls %s", err)
		return
	}
	if errApi != nil {
		t.Errorf("errApi responseMtls : %#v", errApi.Message)
		return
	}

	clientPix := bankly.NewClient(os.Getenv("ENV"))
	clientPix.SetCertificateMtls(certificate)
	responseTokenPix, err := clientPix.RequestToken(responseMtls.ClientID, "", bankly.GetScope().EventsRead, true)
	if err != nil {
		t.Errorf("err : responseTokenPix%s", err)
		return
	}

	clientPix.SetBearer(responseTokenPix.AccessToken)
	clientPix.SetCertificateMtls(certificate)
	begin, err := time.Parse(time.RFC3339, "2022-06-22T00:00:01+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	end, err := time.Parse(time.RFC3339, "2022-06-22T23:59:01+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	response, errApi, err := clientPix.Event().Get(&bankly.EventGetRequest{
		Branch:         "0001",
		Account:        "88192849",
		Page:           "1",
		PageSize:       "100",
		IncludeDetails: true,
		CorrelationID:  "",
		BeginDateTime:  &begin,
		EndDateTime:    &end,
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}
	t.Error("response is null")
	if response == nil {
		t.Error("response is null")
		return
	}

	for _, item := range *response {
		fmt.Printf("Amount : \n%f\n", item.Amount)
		fmt.Printf("Amount : \n%f\n", item.Amount)
	}
}
