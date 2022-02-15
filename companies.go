package snipeit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllCompanies(authToken *string) (*[]Company, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/companies", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	companies := []Company{}
	err = json.Unmarshal(body, &companies)
	if err != nil {
		return nil, err
	}

	return &companies, nil
}
