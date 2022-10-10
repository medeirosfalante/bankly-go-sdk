package bankly

import (
	"encoding/json"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Account is a structure manager all about account
type Pix struct {
	client *Bankly
}

type PixKeys []*PixKey
type PixKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PixCreateKeyRequest struct {
	AddressingKey *PixKey        `json:"addressingKey"`
	Account       *PixKeyAccount `json:"account"`
}

type PixKeyResponse struct {
	EndToEndId    string         `json:"endToEndId"`
	AddressingKey *PixKey        `json:"addressingKey"`
	Account       *PixKeyAccount `json:"account"`
	Status        string         `json:"status"`
	CreatedAt     *time.Time     `json:"createdAt"`
	OwnedAt       *time.Time     `json:"ownedAt"`
	Holder        *PixKeyHolder  `json:"holder,omitempty"`
}

type PixKeyHolder struct {
	Type           string `json:"type"`
	SocialName     string `json:"socialName"`
	TradingName    string `json:"tradingName"`
	DocumentNumber string `json:"documentNumber"`
	Name           string `json:"name"`
}

type PixKeyAccount struct {
	Branch string        `json:"branch"`
	Number string        `json:"number"`
	Type   string        `json:"type"`
	Bank   *Bank         `json:"bank,omitempty"`
	Holder *PixKeyHolder `json:"holder,omitempty"`
}

type PixKeyCashOutRequest struct {
	Sender             *PixCashOutPeople `json:"sender"`
	Recipient          *PixCashOutPeople `json:"recipient"`
	Amount             float64           `json:"amount"`
	Description        string            `json:"description"`
	EndToEndId         string            `json:"endToEndId"`
	AddressKey         string            `json:"addressKey"`
	ConciliationId     string            `json:"conciliationId"`
	InitializationType string            `json:"initializationType,omitempty"`
	DocumentNumber     string            `json:"documentNumber"`
	Name               string            `json:"name"`
}

type PixKeyCashOutResponse struct {
	Sender             *PixCashOutPeople `json:"sender"`
	Recipient          *PixCashOutPeople `json:"recipient"`
	Amount             float64           `json:"amount"`
	Description        string            `json:"description"`
	EndToEndId         string            `json:"endToEndId"`
	AuthenticationCode string            `json:"authenticationCode"`
}

type BankPix struct {
	Name  string `json:"name"`
	Compe string `json:"compe"`
	Ispb  string `json:"ispb"`
}

type PixKeyAccountPeople struct {
	Branch string `json:"branch"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

type PixCashOutPeople struct {
	HolderType     string               `json:"holderType,omitempty"`
	Account        *PixKeyAccountPeople `json:"account"`
	DocumentNumber string               `json:"documentNumber"`
	Name           string               `json:"name"`
	Bank           *BankPix             `json:"bank"`
}

type PixCashOutGet struct {
	Account            string `json:"account"`
	AuthenticationCode string `json:"-"`
}

type BrcodeRequest struct {
	AddressingKey  *PixKey   `json:"addressingKey"`
	Amount         float32   `json:"amount,omitempty"`
	ConciliationId string    `json:"conciliationId,omitempty"`
	CategoryCode   string    `json:"categoryCode,omitempty"`
	AdditionalData string    `json:"additionalData,omitempty"`
	Location       *Location `json:"location"`
	RecipientName  string    `json:"recipientName"`
	OwnerDocument  string
}

type Location struct {
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

type BrcodeResponse struct {
	EncodedValue string `json:"encodedValue"`
}

type GetBercodeResponse struct {
	EncodedValue  string `json:"encodedValue"`
	OwnerDocument string
}

// Pix - Instance de Pix
func (c *Bankly) Pix() *Pix {
	return &Pix{client: c}
}

func (a *Pix) CreateCashOut(req *PixKeyCashOutRequest) (*PixKeyCashOutResponse, *Error, error) {
	uuid := uuid.NewV4()
	var response *PixKeyCashOutResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.RequestPix("POST", "pix/cash-out", uuid.String(), "", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) GetKeys(accountNumber string) (*PixKeys, *Error, error) {
	uuid := uuid.NewV4()
	var response *PixKeys
	err, errApi := a.client.RequestPix("GET", fmt.Sprintf("accounts/%s/addressing-keys", accountNumber), uuid.String(), "", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) CreateKey(req *PixCreateKeyRequest) (*PixKeyResponse, *Error, error) {
	var response *PixKeyResponse
	uuid := uuid.NewV4()
	data, _ := json.Marshal(req)
	err, errApi := a.client.RequestPix("POST", "baas/pix/entries", uuid.String(), "", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) DeleteKey(key string) (bool, *Error, error) {
	uuid := uuid.NewV4()
	err, errApi := a.client.RequestPix("DELETE", fmt.Sprintf("pix/entries/%s", key), uuid.String(), "", nil, nil)
	if err != nil {
		return false, nil, err
	}
	if errApi != nil {
		return false, errApi, nil
	}
	return true, nil, nil
}

func (a *Pix) GetKey(addressingKeyValue string, ownerDocument string) (*PixKeyResponse, *Error, error) {
	var response *PixKeyResponse
	uuid := uuid.NewV4()
	err, errApi := a.client.RequestPix("GET", fmt.Sprintf("pix/entries/%s", addressingKeyValue), uuid.String(), ownerDocument, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) Get(req *PixCashOutGet) (*TransferResponse, *Error, error) {
	var response *TransferResponse
	err, errApi := a.client.Request("GET", fmt.Sprintf("pix/cash-out/accounts/%s/authenticationcode/%s", req.Account, req.AuthenticationCode), "", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) CreateBrcode(req *BrcodeRequest) (*BrcodeResponse, *Error, error) {
	var response *BrcodeResponse
	uuid := uuid.NewV4()
	data, _ := json.Marshal(req)
	err, errApi := a.client.RequestPix("POST", "pix/qrcodes", uuid.String(), req.OwnerDocument, data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) GetBrCode(req *GetBercodeResponse) (*PixKeyResponse, *Error, error) {
	var response *PixKeyResponse
	uuid := uuid.NewV4()
	data, _ := json.Marshal(req)
	err, errApi := a.client.RequestPix("POST", "pix/qrcodes/decode", uuid.String(), req.OwnerDocument, data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
