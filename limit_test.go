package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetLimit(t *testing.T) {
	godotenv.Load(".env.test")
	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().LimitsRead, false)
	if err != nil {
		t.Errorf("err token : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	response, errApi, err := client.Limit().Get(&bankly.LimitGet{
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
