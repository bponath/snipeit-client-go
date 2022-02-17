package snipeit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetOrder - Returns a specifc order
func (c *Client) GetConsumable(consumableID string, authToken *string) (*Consumable, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/consumables/%s", c.HostURL, consumableID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	consumable := Consumable{}
	err = json.Unmarshal(body, &consumable)
	if err != nil {
		return nil, err
	}

	return &consumable, nil
}
