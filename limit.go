package bankly

import (
	"fmt"
	"net/url"
)

// Limit is a structure manager all about bankslip
type Limit struct {
	client *Bankly
}

type LimitGet struct {
	DocumentNumber string `json:"document_number"`
	FeatureName    string `json:"feature_name"`
	CycleType      string `json:"cycle_type"`
	LevelType      string `json:"level_type"`
}

type LimitResponse struct {
	DocumentNumber string  `json:"document_number"`
	FeatureName    string  `json:"feature_name"`
	CycleType      string  `json:"cycle_type"`
	LevelType      string  `json:"level_type"`
	Amount         float64 `json:"amount"`
}

//Limit - Instance de bankslip
func (c *Bankly) Limit() *Limit {
	return &Limit{client: c}
}

func (a *Limit) Get(req *LimitGet) (*LimitResponse, *Error, error) {
	var response *LimitResponse
	params := url.Values{}
	params.Add("featureName", req.FeatureName)
	params.Add("levelType", req.LevelType)
	params.Add("cycleType", req.CycleType)

	err, errApi := a.client.Request("GET", fmt.Sprintf("holders/%s/max-limits?%s", req.DocumentNumber, params.Encode()), "", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
