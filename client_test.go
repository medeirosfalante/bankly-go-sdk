package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestRequesttoken(t *testing.T) {
	godotenv.Load(".env.test")
	client := bankly.NewClient(os.Getenv("ENV"))

	response, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().AccountRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(response.AccessToken) <= 0 {
		t.Errorf("AccessToken is invalid")
	}
	client.SetBearer(response.TokenType)

}
