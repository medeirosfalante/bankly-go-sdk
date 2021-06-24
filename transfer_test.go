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
	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
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
