package bankly

import (
	"encoding/json"
	"time"
)

// Bankslip is a structure manager all about bankslip
type Bankslip struct {
	client *Bankly
}

type BankslipRequest struct {
	Alias          string           `json:"alias"`
	Account        *BankslipAccount `json:"account"`
	DocumentNumber string           `json:"documentNumber"`
	Amount         float32          `json:"amount"`
	DueDate        time.Time        `json:"dueDate"`
}

type BankslipAccount struct {
	Number string `json:"number"`
	Branch string `json:"branch"`
}

type BankslipResponse struct {
	AuthenticationCode string           `json:"authenticationCode"`
	Account            *BankslipAccount `json:"account"`
	UpdatedAt          time.Time        `json:"updatedAt"`
	OurNumber          string           `son:"ourNumber"`
	Aigitable          string           `json:"digitable"`
	Status             string           `json:"status"`
	Amount             BankslipAmount   `json:"amount"`
	DueDate            time.Time        `json:"dueDate"`
}

type BankslipAmount struct {
	Value float64 `json:"value"`
}

type BankslipRequestGet struct {
}

//Bankslip - Instance de bankslip
func (c *Bankly) Bankslip() *Bankslip {
	return &Bankslip{client: c}
}

func (a *Bankslip) Create(req *BankslipRequest) (*BankslipResponse, *Error, error) {
	var response *BankslipResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "bankslip", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Bankslip) Get(req *BankslipRequest) (*BankslipResponse, *Error, error) {
	var response *BankslipResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "bankslip", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
