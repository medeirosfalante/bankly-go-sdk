package bankly

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strings"
	"time"
)

type Scope struct {
	CardCreate             string
	CardUpdate             string
	CardRead               string
	CardPciPassswordUpdate string
	AccountCreate          string
	AccountRead            string
	AccountClose           string
	CustomerWrite          string
	CustomerRead           string
	CustomerCancel         string
	BusinessWrite          string
	BusinessRead           string
	BusinessCancel         string
	KycDocumentWrite       string
	KycDocumentRead        string
	IncomeReportRead       string
	BoletoCreate           string
	BoletoRead             string
	BoletoDelete           string
	LimitsWrite            string
	LimitsRead             string
	PaymentValidate        string
	PaymentConfirm         string
	PaymentRead            string
	TedCashoutCreate       string
	TedCashoutRead         string
	PixAccountRead         string
	PixEntriesCreate       string
	PixEntriesDelete       string
	PixEntriesRead         string
	PixQrcodeCreate        string
	PixQrcodeRead          string
	PixCashoutCreate       string
	PixCashoutRead         string
	EventsRead             string
}

func GetScope() Scope {

	return Scope{
		CardCreate:             "card.create",
		CardUpdate:             "card.update",
		CardRead:               "card.read",
		CardPciPassswordUpdate: "card.pci.password.update",
		AccountCreate:          "account.create",
		AccountRead:            "account.read",
		AccountClose:           "account.close",
		CustomerWrite:          "customer.write",
		CustomerRead:           "customer.read",
		CustomerCancel:         "customer.cancel",
		BusinessWrite:          "business.write",
		BusinessRead:           "business.read",
		BusinessCancel:         "business.cancel",
		KycDocumentWrite:       "kyc.document.write",
		KycDocumentRead:        "kyc.document.read",
		IncomeReportRead:       "income.report.read",
		BoletoCreate:           "boleto.create",
		BoletoRead:             "boleto.read",
		BoletoDelete:           "boleto.delete",
		LimitsWrite:            "limits.write",
		LimitsRead:             "limits.read",
		PaymentValidate:        "payment.validate",
		PaymentConfirm:         "payment.confirm",
		PaymentRead:            "payment.read",
		TedCashoutCreate:       "ted.cashout.create",
		TedCashoutRead:         "ted.cashout.read",
		PixAccountRead:         "pix.account.read",
		PixEntriesCreate:       "pix.entries.create",
		PixEntriesDelete:       "pix.entries.delete",
		PixEntriesRead:         "pix.entries.read",
		PixQrcodeCreate:        "pix.qrcode.create",
		PixQrcodeRead:          "pix.qrcode.read",
		PixCashoutCreate:       "pix.cashout.create",
		PixCashoutRead:         "pix.cashout.read",
		EventsRead:             "events.read",
	}

}

type Bankly struct {
	Client     *http.Client
	Env        string
	Token      string
	ApiVersion string
	Boundary   string
	Scope      string
}

type Error struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Body      string `json:"body"`
}

type TokenRequest struct {
	GrantType    string `json:"grant_type	"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Client      *Bankly
}

func NewClient(env string) *Bankly {
	bankly := &Bankly{
		Client:     &http.Client{Timeout: 60 * time.Second},
		Env:        env,
		ApiVersion: "1.0",
		Boundary:   "---011000010111000001101001",
	}
	return bankly

}

func (bankly *Bankly) RequestFile(method, action, filepathRef, documentType, documentSide string, out interface{}) (error, *Error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.SetBoundary(bankly.Boundary)
	_ = writer.WriteField("documentType", documentType)
	_ = writer.WriteField("documentSide", documentSide)

	file, err := os.Open(filepathRef)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	contentType, err := bankly.GetFileContentType(file)
	if err != nil {
		return err, nil
	}

	buff, err := ioutil.ReadFile(filepathRef)
	if err != nil {
		return err, nil
	}

	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", contentType)
	mh.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"image\"; filename=\"%s\"", file.Name()))
	part3, err := writer.CreatePart(mh)
	if err != nil {
		return err, nil
	}

	part3.Write(buff)

	err = writer.Close()
	if err != nil {
		return err, nil
	}

	url := bankly.devProd()

	endpoint := fmt.Sprintf("%s/%s", url, action)

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return err, nil
	}

	req.Header.Add("api-version", bankly.ApiVersion)
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bankly.Token))
	res, err := bankly.Client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 202 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		errAPI.Body = string(bodyResponse)
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

func (bankly *Bankly) RequestGetFile(action string) ([]byte, error, *Error) {
	url := bankly.devProd()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bankly.Token))
	req.Header.Add("api-version", bankly.ApiVersion)
	return bankly.RequestMaster(req, nil)
}

func (bankly *Bankly) Request(method, action, correlationID string, body []byte, out interface{}) (error, *Error) {
	url := bankly.devProd()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bankly.Token))
	req.Header.Add("api-version", bankly.ApiVersion)
	if correlationID != "" {
		req.Header.Add("x-correlation-id", correlationID)
	}
	_, err, errBody := bankly.RequestMaster(req, &out)
	return err, errBody
}

func (bankly *Bankly) RequestPix(method, action, correlationID, xBklyPixUserId string, body []byte, out interface{}) (error, *Error) {
	url := bankly.devProd()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	if xBklyPixUserId != "" {
		req.Header.Add("x-bkly-pix-user-id", xBklyPixUserId)
	}
	if correlationID != "" {
		req.Header.Add("x-correlation-id", correlationID)
	}
	_, err, errBody := bankly.RequestMaster(req, &out)
	return err, errBody
}

func (bankly *Bankly) RequestMtls(method, action, xBklyPixUserId string, body []byte, out interface{}) (error, *Error) {
	url := bankly.devProdMtls()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}
	if xBklyPixUserId != "" {
		req.Header.Add("x-bkly-pix-user-id", xBklyPixUserId)
	}
	_, err, errBody := bankly.RequestMaster(req, &out)
	return err, errBody
}

func (bankly *Bankly) RequestMaster(req *http.Request, out interface{}) ([]byte, error, *Error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bankly.Token))
	req.Header.Add("api-version", bankly.ApiVersion)
	res, err := bankly.Client.Do(req)
	if err != nil {
		return nil, err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 202 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return bodyResponse, err, nil
		}
		errAPI.Body = string(bodyResponse)
		return bodyResponse, nil, &errAPI
	}
	if out != nil {
		err = json.Unmarshal(bodyResponse, out)
		if err != nil {
			return bodyResponse, err, nil
		}
	}

	return bodyResponse, nil, nil
}

func (Bankly *Bankly) devProd() string {
	if Bankly.Env == "develop" {
		return "https://api.sandbox.bankly.com.br"
	}
	return "https://api.bankly.com.br"
}

func (Bankly *Bankly) TokenUri() string {
	if Bankly.Env == "develop" {
		return "https://login.sandbox.bankly.com.br/connect/token"
	}
	return "https://login.bankly.com.br/connect/token"
}

func (Bankly *Bankly) TokenUriMTls() string {
	if Bankly.Env == "develop" {
		return "https://auth-mtls.sandbox.bankly.com.br/oauth2/token"
	}
	return "https://auth-mtls.bankly.com.br/oauth2/token"
}

func (bankly *Bankly) RequestToken(clientID, clientSecret, scope string, mtls bool) (*TokenResponse, error) {
	var tokenResponse TokenResponse
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", clientID)
	if clientSecret != "" {
		params.Add("client_secret", clientSecret)
	}
	if scope != "" {
		params.Add("scope", scope)
	}

	urlRef := bankly.TokenUri()
	if mtls {
		urlRef = bankly.TokenUriMTls()
	}

	req, err := http.NewRequest("POST", urlRef, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := bankly.Client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	fmt.Printf("bodyResponse \n%s\n ", string(bodyResponse))
	if res.StatusCode > 202 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyResponse))
	}
	err = json.Unmarshal(bodyResponse, &tokenResponse)
	if err != nil {
		return nil, err
	}
	bankly.Token = tokenResponse.AccessToken
	return &tokenResponse, nil
}

func (bankly *Bankly) SetBearer(token string) {
	bankly.Token = token
}

func (bankly *Bankly) SetCertificateMtls(certificate tls.Certificate) *Bankly {
	bankly.Client = &http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{certificate},
			},
		},
	}
	return bankly
}

func (bankly *Bankly) GetFileContentType(out *os.File) (string, error) {

	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func (Bankly *Bankly) devProdMtls() string {
	if Bankly.Env == "develop" {
		return "https://auth-mtls.sandbox.bankly.com.br"
	}
	return "https://auth-mtls.bankly.com.br"
}
