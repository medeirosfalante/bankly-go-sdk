package bankly_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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
		Document:     "66896639652",
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
	response, errApi, err := client.Account().GetAnalysis("66896639652", []string{TokenSend})
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

func TestRegisterClient(t *testing.T) {
	godotenv.Load(".env.test")

	str := "1991-02-02T11:45:26.371Z"
	birthDate, err := time.Parse(time.RFC3339, str)

	client := bankly.NewClient(os.Getenv("BANKLY_CLIENT_ID"), os.Getenv("BANKLY_CLIENT_SECRET"), os.Getenv("ENV"))
	errApi, err := client.Account().RegisterClient(&bankly.AccountClient{
		Phone: &bankly.AccountClientPhone{
			CountryCode: "55",
			Number:      "85983301776",
		},
		Address: &bankly.AccountClientAddress{
			ZipCode:        "60766075",
			AddressLine:    "Vila Francisca Amélia 801",
			Complement:     "",
			Neighborhood:   "Planalto Ayrton Senna",
			City:           "Fortaleza",
			State:          "CE",
			Country:        "BR",
			BuildingNumber: "801",
		},
		SocialName:   "Erick",
		RegisterName: "Erick Cláudio Marcos Vinicius da Costa",
		BirthDate:    birthDate,
		MotherName:   "Marina Raquel",
		Email:        "aagathajoanacavalcanti@sabereler.com.br",
		Document:     "66896639652",
	})
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errApi != nil {
		t.Errorf("errApi : %#v", errApi)
		return
	}

}
