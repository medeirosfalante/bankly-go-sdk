package bankly

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Transfer is a structure manager all about bankslip
type Transfer struct {
	client *Bankly
}

type TransferRequest struct {
	Sender        *TransferSender    `json:"sender"`
	Recipient     *TransferRecipient `json:"recipient"`
	Description   string             `json:"description"`
	Amount        int32              `json:"amount"`
	CorrelationID string             `json:"-"`
}

type TransferSender struct {
	Branch   string        `json:"branch"`
	Account  string        `json:"account"`
	Document string        `json:"document"`
	Name     string        `json:"name"`
	Bank     *TransferBank `json:"bank"`
}

type TransferGet struct {
	Branch             string `json:"branch"`
	Account            string `json:"account"`
	CorrelationID      string `json:"-"`
	AuthenticationCode string `json:"-"`
}

type TransferRecipient struct {
	Branch      string        `json:"branch"`
	Account     string        `json:"account"`
	BankCode    string        `json:"bankCode"`
	Document    string        `json:"document"`
	Name        string        `json:"name"`
	AccountType string        `json:"accountType"`
	Bank        *TransferBank `json:"bank"`
}

type TransferBank struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Ispb string `json:"ispb"`
}

type TransferGetSender struct {
	Branch   string `json:"branch"`
	Account  string `json:"account"`
	Document string `json:"document"`
	Name     string `json:"name"`
}

type TransferGetRecipient struct {
	Branch      string `json:"branch"`
	Account     string `json:"account"`
	BankCode    string `json:"bankCode"`
	Document    string `json:"document"`
	Name        string `json:"name"`
	AccountType string `json:"accountType"`
}

type TransferAccount struct {
	Bank    *TransferBank        `json:"bank"`
	Branch  string               `json:"branch"`
	Account *TransferItemAccount `json:"account"`
}

type TransferItemAccount struct {
	Bank    *TransferBank `json:"bank"`
	Branch  string        `json:"branch"`
	Account string        `json:"account"`
}

type TransferResponse struct {
	AuthenticationCode string           `json:"authenticationCode"`
	Channel            string           `json:"channel"`
	Status             string           `json:"status"`
	CreatedAt          time.Time        `json:"createdAt"`
	UpdatedAt          time.Time        `json:"updatedAt"`
	Sender             *TransferAccount `json:"sender"`
	Recipient          *TransferAccount `json:"recipient"`
	Description        string           `json:"description"`
	Amount             float32          `json:"amount"`
	CorrelationID      string           `json:"-"`
}

//Transfer - Instance de bankslip
func (c *Bankly) Transfer() *Transfer {
	return &Transfer{client: c}
}

func (a *Transfer) Create(req *TransferRequest) (*TransferResponse, *Error, error) {
	var response *TransferResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "fund-transfers", req.CorrelationID, data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Transfer) Get(req *TransferGet) (*TransferResponse, *Error, error) {
	var response *TransferResponse
	params := url.Values{}
	params.Add("account", req.Account)
	params.Add("branch", req.Branch)
	err, errApi := a.client.Request("GET", fmt.Sprintf("fund-transfers/%s?%s", req.AuthenticationCode, params.Encode()), req.CorrelationID, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
