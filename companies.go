package snipeit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllCompanies(authToken *string) (*[]CompanyRow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/companies", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	company := Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		return nil, err
	}

	return &company.Rows, nil
}

// GetOrder - Returns a specifc order
func (c *Client) GetCompany(companyID string, authToken *string) (*CompanyRow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/companies/%s", c.HostURL, companyID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	company := CompanyRow{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

// CreateCompany - Create new company
func (c *Client) CreateCompany(company CompanyRow, authToken *string) (*CompanyRow, error) {
	// rb, err := json.Marshal(company)
	company_to_create := make(map[string]interface{})
	company_to_create["name"] = company.Name

	rb, err := json.Marshal(company_to_create)

	// req, err := http.NewRequest("POST", fmt.Sprintf("%s/companies", c.HostURL), strings.NewReader(string(rb)))
	req, err := http.Post(fmt.Sprintf("%s/companies", c.HostURL), "application/json", bytes.NewBuffer(rb))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	company_return := CompanyRow{}
	err = json.Unmarshal(body, &company_return)
	if err != nil {
		return nil, err
	}

	return &company_return, nil
}
