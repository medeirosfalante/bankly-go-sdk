package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransfer(t *testing.T) {
	godotenv.Load(".env.test")
	uuid := uuid.NewV4()
	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"), bankly.GetScope().TedCashoutCreate)
	response, errApi, err := client.Transfer().Create(&bankly.TransferRequest{
		Sender: &bankly.TransferSender{
			Branch:   "0001",
			Account:  "189863",
			Document: "88698145566",
			Name:     "Joao",
		},
		Recipient: &bankly.TransferRecipient{
			Branch:      "0001",
			Account:     "189863",
			Document:    "88698145566",
			BankCode:    "341",
			Name:        "Joao",
			AccountType: "CHECKING",
		},
		Description:   "",
		Amount:        10 * 100,
		CorrelationID: uuid.String(),
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

func TestGetTransfer(t *testing.T) {
	godotenv.Load(".env.test")
	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"), bankly.GetScope().TedCashoutRead)
	response, errApi, err := client.Transfer().Get(&bankly.TransferGet{
		Branch:             "0001",
		Account:            "199265",
		CorrelationID:      "167a0eb3-3cb2-4767-b7f4-be6ea7d4edf4",
		AuthenticationCode: "04ffa878-ad47-4fd1-8019-22632ba245b9",
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
