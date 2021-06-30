package bankly

import (
	"encoding/json"
	"fmt"
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

type PixKeyAccount struct {
	Branch string `json:"branch"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

//Pix - Instance de Pix
func (c *Bankly) Pix() *Pix {
	return &Pix{client: c}
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
