package bankly_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestCreateBankslip(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Bankslip().Create(&bankly.BankslipRequest{
		Alias:          "",
		Account:        &bankly.BankslipAccount{Branch: "0001", Number: "189863"},
		DocumentNumber: "88698145566",
		Amount:         30,
		DueDate:        time.Now().AddDate(0, 0, 1),
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

func TestGetBankslip(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Bankslip().Get(&bankly.BankslipGetRequest{
		Branch:             "0001",
		Number:             "189863",
		AuthenticationCode: "2bed1d37-3ea8-4b06-9dc8-199883eb4609",
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

func TestGetFile(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Bankslip().GetPdf("2bed1d37-3ea8-4b06-9dc8-199883eb4609")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}

	err = ioutil.WriteFile("test.pdf", response, 0644)
	if err != nil {
		t.Errorf("err : %#v", err)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}

}
