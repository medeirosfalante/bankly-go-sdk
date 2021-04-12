package bankly

import (
	"encoding/json"
	"fmt"
	"net/url"
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
