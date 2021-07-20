package bankly

import (
	"fmt"
	"net/url"
)

// Account is a structure manager all about account
type Services struct {
	client *Bankly
}

//Account - Instance de account
func (c *Bankly) Services() *Services {
	return &Services{client: c}
}

type ListBanksRequest struct {
	Product string
	Ids     []string
	Name    string
}

type ListBanksResponse []*ListBanksItem

type ListBanksItem struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	Ispb      string `json:"ispb"`
	ShortName string `json:"shortName"`
}

func (a *Services) ListBanks(req *ListBanksRequest) (ListBanksResponse, *Error, error) {
	var response ListBanksResponse

	params := url.Values{}
	params.Add("product", req.Product)
	for _, item := range req.Ids {
		params.Add("id", item)
	}
	if req.Name != "" {
		params.Add("name", req.Name)
	}
	err, errApi := a.client.Request("GET", fmt.Sprintf("banklist?%s", params.Encode()), "", nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errApi != nil {
		return nil, errApi, nil
	}
	return response, nil, nil
}
