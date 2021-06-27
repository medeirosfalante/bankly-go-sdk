package bankly_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetEvent(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Event().Get(&bankly.EventGetRequest{
		Branch:         "0001",
		Account:        "200514",
		Page:           "1",
		PageSize:       "10",
		IncludeDetails: true,
		CorrelationID:  "",
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
