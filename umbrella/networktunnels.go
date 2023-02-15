package umbrella

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetTunnel - Returns a specifc tunnel
func (c *Client) GetTunnel(tunnelID string, authToken *string) (*NetworkTunnel, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments/v2/tunnels/%s", c.HostURL, tunnelID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tunnel := NetworkTunnel{}
	err = json.Unmarshal(body, &tunnel)
	if err != nil {
		return nil, err
	}

	return &tunnel, nil
}

// CreateTunnel - Create new Tunnel
func (c *Client) CreateTunnel(TunnelItem NetworkTunnelCreate, authToken *string) (*NetworkTunnel, error) {
	rb, err := json.Marshal(TunnelItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/deployments/v2/tunnels", c.HostURL), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	tunnel := NetworkTunnel{}
	err = json.Unmarshal(body, &tunnel)
	if err != nil {
		return nil, err
	}

	return &tunnel, nil
}

// DeleteTunnel - Deletes a tunnel
func (c *Client) DeleteTunnel(tunnelID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/deployments/v2/tunnels/%s", c.HostURL, tunnelID), nil)
	if err != nil {
		return err
	}
	fmt.Println(req)
	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	response := Response{}
	err = json.Unmarshal(body, &response)

	if string(response.Message) != "Tunnel deleted successfully" {
		return errors.New(string(body))
	}

	return nil
}

// UpdateTunnel - Updates a tunnel
func (c *Client) UpdateTunnel(tunnelID string, TunnelItem NetworkTunnel, authToken *string) (*NetworkTunnel, error) {
	rb, err := json.Marshal(TunnelItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/deployments/v2/tunnels/%s", c.HostURL, tunnelID), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	tunnel := NetworkTunnel{}
	err = json.Unmarshal(body, &tunnel)
	if err != nil {
		return nil, err
	}

	return &tunnel, nil
}