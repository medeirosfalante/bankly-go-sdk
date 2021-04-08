package bankly_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/medeirosfalante/bankly-go-sdk"
)

var TokenSend = ""

func TestSendSelfDocument(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	response, errApi, err := client.Account().SendDocument(&bankly.AccountDocumentRequest{
		DocumentType: "SELFIE",
		DocumentSide: "FRONT",
		ImagePath:    fmt.Sprintf("%s/image-example.jpg", dir),
		Document:     "03602763501",
		DocumentName: "image-example.jpg",
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

	TokenSend = response.Token

}

func TestCheckTokenDocument(t *testing.T) {
	godotenv.Load(".env.test")

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	response, errApi, err := client.Account().GetAnalysis("03602763501", []string{TokenSend})
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
