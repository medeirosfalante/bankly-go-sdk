package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetPaylBill(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))

	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().PixQrcodeRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.PayBill().Validate(&bankly.ValidateBillRequest{
		Code:          "",
		CorrelationID: "233e8bcd-b641-4448-8bf2-9b5288a1d5ad",
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
