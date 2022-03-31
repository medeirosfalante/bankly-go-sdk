package bankly_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

func TestGetEvent(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("ENV"))
	responseToken, err := client.RequestToken(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), bankly.GetScope().EventsRead, false)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	client.SetBearer(responseToken.AccessToken)
	begin, err := time.Parse(time.RFC3339, "2022-03-29T12:00:01+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	end, err := time.Parse(time.RFC3339, "2022-03-29T23:00:01+00:00")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	response, errApi, err := client.Event().Get(&bankly.EventGetRequest{
		Branch:         "0001",
		Account:        "44409281",
		Page:           "1",
		PageSize:       "100",
		IncludeDetails: true,
		CorrelationID:  "",
		BeginDateTime:  &begin,
		EndDateTime:    &end,
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}
	t.Error("response is null")
	if response == nil {
		t.Error("response is null")
		return
	}

	for _, item := range *response {
		fmt.Printf("Amount : \n%f\n", item.Amount)
	}
}
