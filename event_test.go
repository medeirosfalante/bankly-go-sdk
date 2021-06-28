package bankly_test

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetEvent(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	begin, err := time.Parse(time.RFC3339, "2021-06-26T03:26:01+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	end, err := time.Parse(time.RFC3339, "2021-06-26T03:26:08+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	response, errApi, err := client.Event().Get(&bankly.EventGetRequest{
		Branch:         "0001",
		Account:        "200514",
		Page:           "1",
		PageSize:       "10",
		IncludeDetails: true,
		CorrelationID:  "",
		BeginDateTime:  &begin,
		EndDateTime:    &end,
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	t.Errorf("response %#v", response)

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}

	if response == nil {
		t.Error("response is null")
		return
	}
}
