package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestListBanks(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixQrcodeRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Services().ListBanks(&bankly.ListBanksRequest{
		Product: "PIX",
		Ids:     []string{"13140088"},
		Name:    "",
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
