package bankly

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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
}

type PixKeyCashOutResponse struct {
	Sender             *PixCashOutPeople `json:"sender"`
	Recipient          *PixCashOutPeople `json:"recipient"`
	Amount             float64           `json:"amount"`
	Description        string            `json:"description"`
	EndToEndId         string            `json:"endToEndId"`
	AuthenticationCode string            `json:"authenticationCode"`
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
	Bank           *Bank                `json:"bank"`
}

//Pix - Instance de Pix
func (c *Bankly) Pix() *Pix {
	return &Pix{client: c}
}

func (a *Pix) CreateCashOut(req *PixKeyCashOutRequest) (*PixKeyCashOutResponse, *Error, error) {
	var response *PixKeyCashOutResponse
	data, _ := json.Marshal(req)
	log.Printf(" \n%s\n", data)
	err, errApi := a.client.RequestPix("POST", "baas/pix/cash-out", "", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) GetKeys(accountNumber string) (*PixKeys, *Error, error) {
	var response *PixKeys
	err, errApi := a.client.RequestPix("GET", fmt.Sprintf("baas/accounts/%s/addressing-keys", accountNumber), "", nil, &response)
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
	data, _ := json.Marshal(req)
	err, errApi := a.client.RequestPix("POST", "baas/pix/entries", "", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Pix) GetKey(addressingKeyValue string, ownerDocument string) (*PixKeyResponse, *Error, error) {
	var response *PixKeyResponse
	err, errApi := a.client.RequestPix("GET", fmt.Sprintf("baas/pix/entries/%s", addressingKeyValue), ownerDocument, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
