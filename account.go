package bankly

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
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

type AccountAnalysis struct {
	Token           string                          `json:"token"`
	DocumentType    string                          `json:"documentType"`
	DocumentSide    string                          `json:"documentSide"`
	Status          string                          `json:"status"`
	FaceMatch       *AccountAnalysisFaceMatch       `json:"faceMatch"`
	FaceDetails     *AccountAnalysisFaceDetails     `json:"faceDetails"`
	DocumentDetails *AccountAnalysisDocumentDetails `json:"documentDetails"`
	Liveness        *AccountAnalysisLiveness        `json:"liveness"`
	AnalyzedAt      string                          `json:"analyzedAt"`
}

type AccountAnalysisResponse []*AccountAnalysis

type AccountAnalysisDocumentDetails struct {
	Status                          string `json:"Status"`
	IdentifiedDocumentTypestring    string `json:"IdentifiedDocumentTypestring"`
	IdNumber                        string `json:"idNumber"`
	CpfNumber                       string `json:"cpfNumber"`
	BirthDate                       string `json:"birthDate"`
	FatherName                      string `json:"fatherName"`
	MotherName                      string `json:"motherName"`
	RegisterName                    string `json:"registerName"`
	ValidDate                       string `json:"validDate"`
	DriveLicenseCategory            string `json:"driveLicenseCategory"`
	DriveLicenseNumber              string `json:"driveLicenseNumber"`
	DriveLicenseFirstQualifyingDate string `json:"driveLicenseFirstQualifyingDate"`
	FederativeUnit                  string `json:"federativeUnit"`
	IssuedBy                        string `json:"issuedBy"`
	IssuePlace                      string `json:"issuePlace"`
	IssueDate                       string `json:"issueDate"`
}

type AccountClient struct {
	Phone        *AccountClientPhone   `json:"phone"`
	Address      *AccountClientAddress `json:"address"`
	SocialName   string                `json:"socialName"`
	RegisterName string                `json:"registerName"`
	BirthDate    time.Time             `json:"birthDate"`
	MotherName   string                `json:"motherName"`
	Email        string                `json:"email"`
	Document     string
}

type AccountBussiness struct {
	Phone               *AccountClientPhone         `json:"phone"`
	BusinessAddress     *AccountClientAddress       `json:"businessAddress"`
	MotherName          string                      `json:"motherName"`
	BusinessEmail       string                      `json:"businessEmail"`
	BusinessName        string                      `json:"businessName"`
	TradingName         string                      `json:"tradingName"`
	BusinessType        string                      `json:"businessType"`
	BusinessSize        string                      `json:"businessSize"`
	LegalRepresentative *AccountLegalRepresentative `json:"legalRepresentative"`
	Document            string
}

type AccountLegalRepresentative struct {
	Phone        *AccountClientPhone   `json:"phone"`
	Address      *AccountClientAddress `json:"address"`
	SocialName   string                `json:"socialName"`
	RegisterName string                `json:"registerName"`
	BirthDate    time.Time             `json:"birthDate"`
	MotherName   string                `json:"motherName"`
	Email        string                `json:"email"`
}

type AccountClientPhone struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

type AccountClientAddress struct {
	ZipCode        string `json:"zipCode"`
	AddressLine    string `json:"addressLine"`
	Complement     string `json:"complement"`
	Neighborhood   string `json:"neighborhood"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	BuildingNumber string `json:"buildingNumber"`
}

type AccountAnalysisFaceMatch struct {
	Status     string `json:"status"`
	Similarity string `json:"similarity"`
	Confidence string `json:"confidence"`
}

type AccountAnalysisFaceDetails struct {
	Status string `json:"status"`
}
type AccountAnalysisLiveness struct {
}

type RequestNewAccount struct {
	AccountType    string `json:"accountType"`
	DocumentNumber string `json:"-"`
}

type AcountPay struct {
	Balance *Balance `json:"balance"`
	Status  string   `json:"status"`
	Branch  string   `json:"branch"`
	Number  string   `json:"number"`
}

type Balance struct {
	InProcess *BalanceItem `json:"inProcess"`
	Available *BalanceItem `json:"available"`
	Blocked   *BalanceItem `json:"blocked"`
}

type BalanceItem struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type StatementRequest struct {
	Branch  string `json:"branch"`
	Account string `json:"account"`
	Offset  string `json:"offset"`
	Limit   string `json:"limit"`
	Details string `json:"details"`
}

type StatementResponse struct {
	TotalItens int32            `json:"totalItens"`
	Itens      []*StatementItem `json:"itens"`
	PageIndex  int32            `json:"pageIndex"`
}

type StatementItem struct {
	Type         string  `json:"type"`
	Amount       float32 `json:"amount"`
	Operation    string  `json:"operation"`
	CreationDate string  `json:"creationDate"`
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

func (a *Account) GetAnalysis(document string, tokens []string) (*AccountAnalysisResponse, *Error, error) {
	var response *AccountAnalysisResponse
	params := url.Values{}
	for _, item := range tokens {
		params.Add("token", item)
	}
	err, errApi := a.client.Request("GET", fmt.Sprintf("document-analysis/%s?%s", document, params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) RegisterClient(req *AccountClient) (*Error, error) {
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("PUT", fmt.Sprintf("customers/%s", req.Document), data, nil)
	if err != nil {
		return nil, err
	}
	if errApi != nil {
		return errApi, nil
	}
	return nil, nil
}

func (a *Account) GetClient(document string, resultLevel string) (*AccountClient, *Error, error) {
	var response *AccountClient
	params := url.Values{}
	params.Add("resultLevel", resultLevel)
	err, errApi := a.client.Request("GET", fmt.Sprintf("customers/%s?%s", document, params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) RegisterAccount(req *RequestNewAccount) (*AcountPay, *Error, error) {
	var response *AcountPay
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", fmt.Sprintf("customers/%s/accounts", req.DocumentNumber), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) GetAccounts(document string) ([]*AcountPay, *Error, error) {
	var response []*AcountPay
	err, errApi := a.client.Request("GET", fmt.Sprintf("customers/%s/accounts", document), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) GetAccount(account string, includeBalance bool) (*AcountPay, *Error, error) {
	var response *AcountPay
	params := url.Values{}
	params.Add("includeBalance", strconv.FormatBool(includeBalance))
	err, errApi := a.client.Request("GET", fmt.Sprintf("accounts/%s?%s", account, params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

// https://api.sandbox.bankly.com.br/accounts/accountNumber

func (a *Account) GetStatement(req *StatementRequest) (*StatementResponse, *Error, error) {
	params := url.Values{}
	params.Add("account", req.Account)
	params.Add("branch", req.Branch)
	params.Add("limit", req.Limit)
	params.Add("offset", req.Offset)
	params.Add("details", req.Details)
	var response *StatementResponse
	err, errApi := a.client.Request("GET", fmt.Sprintf("account/statement?%s", params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) RegisterBusiness(req *AccountBussiness) (*Error, error) {
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("PUT", fmt.Sprintf("business/%s", req.Document), data, nil)
	if err != nil {
		return nil, err
	}
	if errApi != nil {
		return errApi, nil
	}
	return nil, nil
}

func (a *Account) GetBusiness(document string, resultLevel string) (*AccountBussiness, *Error, error) {
	var response *AccountBussiness
	params := url.Values{}
	params.Add("resultLevel", resultLevel)
	err, errApi := a.client.Request("GET", fmt.Sprintf("business/%s?%s", document, params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) RegisterAccountBusiness(req *RequestNewAccount) (*AcountPay, *Error, error) {
	var response *AcountPay
	data, _ := json.Marshal(req)
	err, errApi := a.client.Request("POST", fmt.Sprintf("business/%s/accounts", req.DocumentNumber), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}

func (a *Account) GetAccountBusiness(account string, includeBalance bool) (*AcountPay, *Error, error) {
	var response *AcountPay
	params := url.Values{}
	params.Add("includeBalance", strconv.FormatBool(includeBalance))
	err, errApi := a.client.Request("GET", fmt.Sprintf("business/accounts/%s?%s", account, params.Encode()), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
