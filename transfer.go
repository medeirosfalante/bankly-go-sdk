package bankly

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// Transfer is a structure manager all about bankslip
type Transfer struct {
	client *Bankly
}

type TransferRequest struct {
	Sender      *TransferSender    `json:"sender"`
	Recipient   *TransferRecipient `json:"recipient"`
	Description string             `json:"description"`
	Amount      int32              `json:"amount"`
}

type TransferSender struct {
	Branch   string `json:"branch"`
	Account  string `json:"account"`
	Document string `json:"document"`
	Name     string `json:"name"`
}

type TransferRecipient struct {
	Branch      string `json:"branch"`
	Account     string `json:"account"`
	BankCode    string `json:"bankCode"`
	Document    string `json:"document"`
	Name        string `json:"name"`
	AccountType string `json:"accountType"`
}

type TransferResponse struct {
	AuthenticationCode string `json:"authenticationCode"`
}

//Transfer - Instance de bankslip
func (c *Bankly) Transfer() *Transfer {
	return &Transfer{client: c}
}

func (a *Transfer) Create(req *TransferRequest) (*TransferResponse, *Error, error) {
	uuid := uuid.NewV4()
	var response *TransferResponse
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", "fund-transfers", uuid.String(), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
