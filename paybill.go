package bankly

import (
	"encoding/json"
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
	Code          string `json:"code"`
	CorrelationID string `json:"correlationId"`
	Description   string `json:"description"`
	BankBranch    string `json:"bankBranch"`
	BankAccount   string `json:"bankAccount"`
	ID            string `json:"id"`
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

func (a *PayBill) Validate(req *ValidateBillRequest) (*ValidateBillResponse, *Error, error) {
	var response *ValidateBillResponse
	err, errApi := a.client.Request("POST", "bill-payment/validate", req.CorrelationID, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
