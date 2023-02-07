package umbrella

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetlistOfNetworks
func (c *Client) GetListNetworks(authToken *string) (*[]Network, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments/v2/networks", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	//trimmedbody := bytes.Trim(body,"[]")
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	var networks []Network

	err = json.Unmarshal(body, &networks)
	if err != nil {
		return nil, err
	}

	return &networks, nil
}

// CreateNetwork - Create new Network
func (c *Client) CreateNetwork(NetworkItem Network, authToken *string) (*Network, error) {
	rb, err := json.Marshal(NetworkItem)
	if err != nil {
		return nil, err
	}

	//rb[2] = byte(unicode.ToLower(rune(rb[2])))

	fmt.Println(string(rb))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/deployments/v2/networks", c.HostURL), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	fmt.Println(req)

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	network := Network{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// CreateInternalNetwork - Create new Internal Network
func (c *Client) CreateInternalNetwork(NetworkItem Internalnetwork, authToken *string) (*Internalnetwork, error) {
	rb, err := json.Marshal(NetworkItem)
	if err != nil {
		return nil, err
	}

	//rb[2] = byte(unicode.ToLower(rune(rb[2])))

	fmt.Println(string(rb))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/deployments/v2/internalnetworks", c.HostURL), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	fmt.Println(req)

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	network := Internalnetwork{}
	err = json.Unmarshal(body, &network)
	if err != nil {
		return nil, err
	}

	return &network, nil
}

// CreateSite - Create new Site
func (c *Client) CreateSite(SiteItem Site, authToken *string) (*Site, error) {
	rb, err := json.Marshal(SiteItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/deployments/v2/sites", c.HostURL), bytes.NewBuffer(rb))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	site := Site{}
	err = json.Unmarshal(body, &site)
	if err != nil {
		return nil, err
	}

	return &site, nil
}

// DeleteSite - Deletes a site
func (c *Client) DeleteSite(siteID string, authToken *string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/deployments/v2/sites/%s", c.HostURL, siteID), nil)
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

// GetSite - Returns a specifc Site
func (c *Client) GetSite(siteID string, authToken *string) (*Site, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/deployments/v2/sites/%s", c.HostURL, siteID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	site := Site{}
	err = json.Unmarshal(body, &site)
	if err != nil {
		return nil, err
	}

	return &site, nil
}
