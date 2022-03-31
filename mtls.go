package bankly

import (
	"encoding/json"
)

// Account is a structure manager all about account
type Mtls struct {
	client *Bankly
}

//Mtls - Instance de Mtls
func (c *Bankly) Mtls() *Mtls {
	return &Mtls{client: c}
}

type RequestRegisterClientID struct {
	GrantTypes              []string `json:"grant_types"`
	TlsClientAuthSubjectDn  string   `json:"tls_client_auth_subject_dn"`
	TokenEndpointAuthMethod string   `json:"token_endpoint_auth_method"`
	ResponseTypes           []string `json:"response_types"`
	CompanyKey              string   `json:"company_key"`
	Scope                   string   `json:"scope"`
}

type ResponseRegisterClientID struct {
	GrantTypes                       []string `json:"grant_types"`
	TlsClientAuthSubjectDn           string   `json:"tls_client_auth_subject_dn"`
	TokenEndpointAuthMethod          string   `json:"token_endpoint_auth_method"`
	ResponseTypes                    []string `json:"response_types"`
	CompanyKey                       string   `json:"company_key"`
	Scope                            string   `json:"scope"`
	SubjectType                      string   `json:"subjectType"`
	RegistrationAccessTokenExpiresIn int64    `json:"registration_access_token_expires_in"`
	RegistrationAccessToken          string   `json:"registration_access_token"`
	TokenEndpointAuthMethods         []string `json:"token_endpoint_auth_methods"`
	ClientID                         string   `json:"client_id"`
	AccessTokenTtl                   int32    `json:"access_token_ttl"`
	ClientIDIssuedAt                 int64    `json:"client_id_issued_at"`
}

func (c *Mtls) RegisterClientID(req *RequestRegisterClientID) (*ResponseRegisterClientID, *Error, error) {
	var response *ResponseRegisterClientID
	data, _ := json.Marshal(req)
	err, errApi := c.client.RequestMtls("POST", "oauth2/register", "", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
