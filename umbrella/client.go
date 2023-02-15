package umbrella

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "https://api.umbrella.com"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Apikey    string
	Apisecret string
}

// AuthResponse -
type AuthResponse struct {
	Tokentype string `json:"token_type"`
	Token     string `json:"access_token"`
	Expiresin int    `json:"expires_in"`
}

// NewClient -
func NewClient(host, apikey, apisecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	if apikey == nil || apisecret == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Apikey:    *apikey,
		Apisecret: *apisecret,
	}

	ar, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
		req.Header.Set("Authorization", "Bearer "+token)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 204 {
		return []byte("204"), nil
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
