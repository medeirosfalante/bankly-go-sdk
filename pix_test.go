package bankly_test

import (
	"crypto/tls"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetKeys(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixAccountRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Pix().GetKeys("199265")
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

func TestCreateKey(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixEntriesCreate, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Pix().CreateKey(&bankly.PixCreateKeyRequest{
		AddressingKey: &bankly.PixKey{
			Type:  "CNPJ",
			Value: "35818953000146",
		},
		Account: &bankly.PixKeyAccount{
			Branch: "0001",
			Number: "199265",
			Type:   "PAYMENT",
		},
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

func TestGetKey(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixEntriesRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Pix().GetKey("35818953000146", "35818953000146")
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

func TestCashOut(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixCashoutCreate, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Pix().CreateCashOut(&bankly.PixKeyCashOutRequest{
		Sender: &bankly.PixCashOutPeople{
			Account: &bankly.PixKeyAccountPeople{
				Branch: "0001",
				Number: "199265",
				Type:   "PAYMENT",
			},
			DocumentNumber: "35818953000146",
			Name:           "GENESIS BANK SOLUCOES DE PAGAMENTO LTDA",
			Bank: &bankly.BankPix{
				Name:  "Acesso Soluções de Pagamento S.A.",
				Compe: "332",
				Ispb:  "13140088",
			},
		},
		Recipient: &bankly.PixCashOutPeople{
			Account: &bankly.PixKeyAccountPeople{
				Branch: "0001",
				Number: "200514",
				Type:   "PAYMENT",
			},
			DocumentNumber: "28503661000159",
			Name:           "ALBERTO ALMEIDA DE AZEVEDO TECNOLOGIA",
			Bank: &bankly.BankPix{
				Name:  "Acesso Soluções de Pagamento S.A.",
				Compe: "332",
				Ispb:  "13140088",
			},
		},
		Amount:      10,
		Description: "",
		EndToEndId:  "",
		AddressKey:  "28503661000159",
		// InitializationType: "KEY",
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

func TestGetPix(t *testing.T) {
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
		Scope:                   bankly.GetScope().PixCashoutRead,
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
	responseTokenPix, err := clientPix.RequestToken(responseMtls.ClientID, "", bankly.GetScope().PixCashoutRead, true)
	if err != nil {
		t.Errorf("err : responseTokenPix%s", err)
		return
	}

	clientPix.SetBearer(responseTokenPix.AccessToken)
	clientPix.SetCertificateMtls(certificate)

	response, errApi, err := clientPix.Pix().Get(&bankly.PixCashOutGet{
		Account:            "88069990",
		AuthenticationCode: "d0540e2f-79a7-438c-a73e-b2d3018705c8",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}
	body, _ := json.Marshal(response)
	t.Errorf("\n\nresponse : %s\n\n", string(body))
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestDeletePix(t *testing.T) {
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
		Scope:                   bankly.GetScope().PixEntriesDelete,
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
	responseTokenPix, err := clientPix.RequestToken(responseMtls.ClientID, "", bankly.GetScope().PixEntriesDelete, true)
	if err != nil {
		t.Errorf("err : responseTokenPix%s", err)
		return
	}

	clientPix.SetBearer(responseTokenPix.AccessToken)
	clientPix.SetCertificateMtls(certificate)

	var list = []string{"6b78f747-463c-4f87-875b-3b6a06ea6bec", "c1384157-994f-48e3-a1ea-2db9e9fa00a1", "5ce078ec-002f-4147-ad6a-274b0c5b766d", "69abd2a3-f7d1-4d72-af1b-32b2fad408ad", "8bec84d9-a79a-4a6c-9b04-1e92ab1419ab", "a9c73e1d-8b18-482e-8384-34b8eb195ba6", "916c5dd7-750e-409e-91a7-2f79dd76151f", "eeae784a-46d4-4bf6-8954-f0fc0492286d"}

	for _, item := range list {
		clientPix.Pix().DeleteKey(item)
		time.Sleep(time.Second * 3)
	}

	// _, errApi, err = clientPix.Pix().DeleteKey("dbb18a88-0ab0-4a5f-8118-b1e52f760acc")
	// if err != nil {
	// 	t.Errorf("err : %s", err)
	// 	return
	// }
	// if errApi != nil {
	// 	t.Errorf("errApi : %#v", errApi)
	// 	return
	// }

}

func TestCreateBrCode(t *testing.T) {
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
		Scope:                   bankly.GetScope().PixQrcodeCreate,
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
	responseTokenPix, err := clientPix.RequestToken(responseMtls.ClientID, "", bankly.GetScope().PixQrcodeCreate, true)
	if err != nil {
		t.Errorf("err : responseTokenPix%s", err)
		return
	}
	clientPix.SetBearer(responseTokenPix.AccessToken)
	clientPix.SetCertificateMtls(certificate)

	response, errApi, err := clientPix.Pix().CreateBrcode(&bankly.BrcodeRequest{
		AddressingKey: &bankly.PixKey{
			Type:  "CNPJ",
			Value: "35818953000146",
		},
		Amount: 2,
		Location: &bankly.Location{
			City:    "SAOPAULO",
			ZipCode: "11111111",
		},
		RecipientName: "JOAO",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errApi != nil {

		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestGetBrCode(t *testing.T) {
	godotenv.Load(".env.test")

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	certificate, err := tls.LoadX509KeyPair(dir+"/cert/client.crt", dir+"/cert/client.key")
	if err != nil {
		t.Errorf("could not load certificate: %v", err)
		return
	}

	client := bankly.NewClient(os.Getenv("ENV"))

	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixQrcodeRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	client.SetCertificateMtls(certificate)
	response, errApi, err := client.Pix().GetBrCode(&bankly.GetBercodeResponse{
		EncodedValue:  "MDAwMjAxMjYzNjAwMTRici5nb3YuYmNiLnBpeDAxMTQzNTgxODk1MzAwMDE0NjUyMDQwMDAwNTMwMzk4NjU0MDQyLjAwNTgwMkJSNTkwOUpPQU8gSk9BTzYwMDhTQU9QQVVMTzYxMDgxMTExMTExMTYyMDcwNTAzKioqNjMwNDNFOUE=",
		OwnerDocument: "35818953000146",
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
