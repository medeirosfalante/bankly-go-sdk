package bankly_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetKeys(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Pix().Get(&bankly.PixCashOutGet{
		Account:            "199265",
		AuthenticationCode: "e8b1eb27-5c6b-45be-9d8d-cc05b4c3a8fc",
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
	t.Errorf("response : %s", string(body))
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestCreateBrCode(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Pix().CreateBrcode(&bankly.BrcodeRequest{
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
		t.Errorf("errApi : %#v", errApi)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestGetBrCode(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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
