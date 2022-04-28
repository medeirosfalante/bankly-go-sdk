package bankly

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// PayBill is a structure manager all about bankslip
type PayBill struct {
	client *Bankly
}

type ValidateBillRequest struct {
	Code          string `json:"code"`
	CorrelationID string `json:"correlationId"`
}

type ValidateBillResponse struct {
	ID                string                     `json:"id"`
	Assignor          string                     `json:"assignor"`
	Recipient         *ValidateBillRecipient     `json:"recipient"`
	Payer             *ValidateBillPayer         `json:"payer"`
	BusinessHours     *ValidateBillBusinessHours `json:"businessHours"`
	DueDate           string                     `json:"dueDate"`
	SettleDate        string                     `json:"settleDate"`
	NextSettle        bool                       `json:"nextSettle"`
	OriginalAmount    float32                    `json:"originalAmount"`
	Amount            float32                    `json:"amount"`
	Charges           *ValidateBillCharges       `json:"charges"`
	MaxAmount         float32                    `json:"maxAmount"`
	MinAmount         float32                    `json:"minAmount"`
	AllowChangeAmount bool                       `json:"allowChangeAmount"`
	Digitable         string                     `json:"digitable"`
}

type ValidateBillRecipient struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
}

type ValidateBillPayer struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
}

type ValidateBillBusinessHours struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"document_number"`
}

type ValidateBillCharges struct {
	InterestAmountCalculated float32 `json:"interestAmountCalculated"`
	FineAmountCalculated     float32 `json:"fineAmountCalculated"`
	DiscountAmount           float32 `json:"discountAmount"`
}

type ConfirmBillRequest struct {
	CorrelationID string  `json:"correlationId"`
	Description   string  `json:"description"`
	BankBranch    string  `json:"bankBranch"`
	BankAccount   string  `json:"bankAccount"`
	ID            string  `json:"id"`
	Amount        float32 `json:"amount"`
}

type GetBillRequest struct {
	AuthenticationCode string `json:"code"`
	CorrelationID      string `json:"correlationId"`
	BankBranch         string `json:"bankBranch"`
	BankAccount        string `json:"bankAccount"`
}

type Bill struct {
	BankBranch         string       `json:"bankBranch"`
	BankAccount        string       `json:"bankAccount"`
	PaymentDate        string       `json:"paymentDate"`
	AuthenticationCode string       `json:"authenticationCode"`
	Status             string       `json:"status"`
	ConfirmedAt        string       `json:"confirmedAt"`
	Digitable          string       `json:"digitable"`
	Amount             float32      `json:"amount"`
	RecipientDocument  string       `json:"recipientDocument"`
	RecipientName      string       `json:"recipientName"`
	SettleDate         string       `json:"settleDate"`
	DueDate            string       `json:"dueDate"`
	Description        string       `json:"description"`
	Charges            *ChargeBills `json:"charges"`
}

type ChargeBills struct {
	InterestAmountCalculated float64 `json:"interestAmountCalculated"`
	FineAmountCalculated     float64 `json:"fineAmountCalculated"`
	DiscountAmount           float64 `json:"discountAmount"`
}

type ConfirmBillResponse struct {
	AuthenticationCode string `json:"authenticationCode"`
	SettleDate         string `json:"settleDate"`
}

//PayBill - Instance de bankslip
func (c *Bankly) PayBill() *PayBill {
	return &PayBill{client: c}
}

func (a *PayBill) Create(req *ConfirmBillRequest) (*ConfirmBillResponse, *Error, error) {
	var response *ConfirmBillResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "bill-payment/confirm", req.CorrelationID, data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *PayBill) Get(req *GetBillRequest) (*Bill, *Error, error) {
	var response *Bill
	params := url.Values{}
	params.Add("bankBranch", req.BankBranch)
	params.Add("bankAccount", req.BankAccount)
	params.Add("authenticationCode", req.AuthenticationCode)
	err, errApi := a.client.Request("GET", fmt.Sprintf("bill-payment/detail?%s", params.Encode()), req.CorrelationID, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *PayBill) Validate(req *ValidateBillRequest) (*ValidateBillResponse, *Error, error) {
	var response *ValidateBillResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "bill-payment/validate", req.CorrelationID, data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
