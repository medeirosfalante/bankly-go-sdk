package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestRequesttoken(t *testing.T) {
	godotenv.Load(".env.test")
	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"), bankly.GetScope().AccountRead)

	response, err := client.RequestToken()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(response.AccessToken) <= 0 {
		t.Errorf("AccessToken is invalid")
	}

}
