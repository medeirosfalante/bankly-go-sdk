package bankly

import (
	"fmt"
)

// Account is a structure manager all about account
type Account struct {
	client *Bankly
}

type AccountDocumentRequest struct {
	DocumentType string `json:"documentType"`
	DocumentSide string `json:"documentSide"`
	ImagePath    string `json:"image"`
	Document     string `json:"document"`
	DocumentName string `json:"document_name"`
}

type AccountDocumentResponse struct {
	Token string `json:"token"`
}

//Account - Instance de account
func (c *Bankly) Account() *Account {
	return &Account{client: c}
}

func (a *Account) SendDocument(req *AccountDocumentRequest) (*AccountDocumentResponse, *Error, error) {
	var response *AccountDocumentResponse
	err, errApi := a.client.RequestFile("PUT", fmt.Sprintf("document-analysis/%s", req.Document), req.ImagePath, req.DocumentType, req.DocumentSide, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
