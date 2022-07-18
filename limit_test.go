package bankly_test

import (
	"crypto/tls"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetLimit(t *testing.T) {
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
		Scope:                   bankly.GetScope().LimitsRead,
	})

	if err != nil {
		t.Errorf("err : responseMtls %s", err)
		return
	}
	if errApi != nil {
		t.Errorf("errApi responseMtls : %#v", errApi.Message)
		return
	}

	clientLimits := bankly.NewClient(os.Getenv("ENV"))
	clientLimits.SetCertificateMtls(certificate)
	responseTokenPix, err := clientLimits.RequestToken(responseMtls.ClientID, "", bankly.GetScope().LimitsRead, true)
	if err != nil {
		t.Errorf("err : responseTokenPix%s", err)
		return
	}
	clientLimits.SetBearer(responseTokenPix.AccessToken)
	clientLimits.SetCertificateMtls(certificate)
	response, errApi, err := clientLimits.Limit().Get(&bankly.LimitGet{
		DocumentNumber: "41246542000126",
		FeatureName:    "SPI",
		CycleType:      bankly.GetLimitType().Monthly,
		LevelType:      "Account",
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
