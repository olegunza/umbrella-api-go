package umbrella

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/clarketm/json"
)

// GetVA - Returns a specific VA
func (c *Client) GetVA(originID int64, authToken *string) (*VA, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments/v2/virtualappliances/%d", c.HostURL, originID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	va := VA{}
	err = json.Unmarshal(body, &va)
	if err != nil {
		return nil, err
	}

	return &va, nil
}

// UpdateVA - Updates a specific VA
func (c *Client) UpdateVA(originID int64, VAItem VA, authToken *string) (*VA, error) {
	rb, err := json.Marshal(VAItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/deployments/v2/virtualappliances/%d", c.HostURL, originID), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	va := VA{}
	err = json.Unmarshal(body, &va)
	if err != nil {
		return nil, err
	}

	return &va, nil
}

// DeleteVA - Deletes a VA
func (c *Client) DeleteVA(originID int64, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/deployments/v2/virtualappliances/%d", c.HostURL, originID), nil)
	if err != nil {
		return err
	}
	fmt.Println(req)
	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "204" {
		return errors.New(string(body))
	}

	return nil
}

// GetVAs - Returns list of virtual appliances
func (c *Client) GetVAs(authToken *string) ([]VA, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments/v2/virtualappliances", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	vas := []VA{}
	err = json.Unmarshal(body, &vas)
	if err != nil {
		return nil, err
	}

	return vas, nil
}
