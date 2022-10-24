package sdk

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const SectigoCertManagerURL = "https://cert-manager.com/api"

type Client struct {
	URL         string
	UserName    string
	CustomerURI string
	Password    string
	HTTPClient  *http.Client
	// Also supports certificate, but that's for later
}

func NewClient(username, customerURI, password *string) *Client {
	c := Client{
		URL:         SectigoCertManagerURL,
		UserName:    *username,
		CustomerURI: *customerURI,
		Password:    *password,
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
	}

	return &c
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("login", c.UserName)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("customerUri", c.CustomerURI)
	req.Header.Set("password", c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
