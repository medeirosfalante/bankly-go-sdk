package bankly

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Bankly struct {
	client       *http.Client
	ClientID     string
	ClientSecret string
	Env          string
	Token        string
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
}

func NewClient(ClientID, ClientSecret, env string) *Bankly {
	return &Bankly{
		client:       &http.Client{Timeout: 60 * time.Second},
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Env:          env,
	}
}

func (bankly *Bankly) Request(method, action string, body []byte, out interface{}) (error, *Error) {
	if bankly.client == nil {
		bankly.client = &http.Client{Timeout: 60 * time.Second}
	}
	url := bankly.devProd()
	endpoint := fmt.Sprintf("%s/%s", url, action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err, nil
	}

	_, err = bankly.RequestToken()
	if err != nil {
		return err, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bankly.Token))
	res, err := bankly.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
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

func (bankly *Bankly) RequestToken() (*TokenResponse, error) {
	var tokenResponse TokenResponse
	if bankly.client == nil {
		bankly.client = &http.Client{Timeout: 60 * time.Second}
	}
	params := url.Values{}
	params.Add("client_secret", bankly.ClientSecret)
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", bankly.ClientID)
	req, err := http.NewRequest("POST", bankly.TokenUri(), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
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
