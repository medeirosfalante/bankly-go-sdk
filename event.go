package bankly

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Transfer is a structure manager all about bankslip
type Event struct {
	client *Bankly
}

type EventGetRequest struct {
	Branch         string     `json:"branch"`
	Account        string     `json:"account"`
	Page           string     `json:"page"`
	PageSize       string     `json:"pageSize"`
	IncludeDetails bool       `json:"includeDetails"`
	BeginDateTime  *time.Time `json:"beginDateTime"`
	EndDateTime    *time.Time `json:"endDateTime"`
	CorrelationID  string     `json:"-"`
}

type EventResponse []*EventItemResponse

type EventItemResponse struct {
	AggregateId    string                 `json:"aggregateId"`
	Type           string                 `json:"type"`
	Category       string                 `json:"category"`
	DocumentNumber string                 `json:"documentNumber"`
	BankBranch     string                 `json:"bankBranch"`
	BankAccount    string                 `json:"bankAccount"`
	Amount         float32                `json:"amount"`
	Name           string                 `json:"name"`
	Timestamp      *time.Time             `json:"timestamp"`
	Data           map[string]interface{} `json:"data"`
}

type EventAccount struct {
	Agency     string `json:"Agency"`
	Account    string `json:"Account"`
	Document   string `json:"Document"`
	IspbNumber string `json:"IspbNumber"`
	Name       string `json:"Name"`
}

type EventDataExternalTransferResponse struct {
	ExternalTransactionId string
	DepositTransactionId  string
	ControlNumber         string
	TransactionAmount     float32
	ClearingAmount        float32
	OverLimitAmount       float32
	Channel               string
	SenderAccount         *EventAccount `json:"SenderAccount"`
	RecipientAccount      *EventAccount `json:"RecipientAccount"`
	CorrelationID         string        `json:"CorrelationId"`
	Document              string        `json:"document"`
	CompanyKey            string        `json:"CompanyKey"`
	ConciliationId        string        `json:"ConciliationId"`
	EndToEndId            string        `json:"EndToEndId"`
	AddressKey            string        `json:"AddressKey"`
}

//Event - Instance de vvent
func (c *Bankly) Event() *Event {
	return &Event{client: c}
}

func (a *Event) Get(req *EventGetRequest) (*EventResponse, *Error, error) {
	var response *EventResponse
	params := url.Values{}
	params.Add("account", req.Account)
	params.Add("branch", req.Branch)
	params.Add("page", req.Page)
	params.Add("pageSize", req.PageSize)
	params.Add("includeDetails", strconv.FormatBool(req.IncludeDetails))
	if req.BeginDateTime != nil {
		params.Add("BeginDateTime", req.BeginDateTime.Format(time.RFC3339))
	}
	if req.EndDateTime != nil {
		params.Add("EndDateTime", req.EndDateTime.Format(time.RFC3339))
	}

	err, errApi := a.client.Request("GET", fmt.Sprintf("events?%s", params.Encode()), req.CorrelationID, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
