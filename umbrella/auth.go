package umbrella

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	_ "strings"
	b64 "encoding/base64"
)

// SignIn - Get a new token for user
func (c *Client) GetToken() (*AuthResponse, error) {
	
	rb := c.Auth.Apikey + ":" + c.Auth.Apisecret
    rbEnc := b64.StdEncoding.EncodeToString([]byte(rb))
	fmt.Println(rbEnc)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/v2/token", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", " Basic "+ rbEnc)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	
	fmt.Println(req)

	body, err := c.doRequest(req, nil)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Unable to login")
	}

	//fmt.Println(string(body))

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}
    
	fmt.Println(ar)
	return &ar, nil
}
