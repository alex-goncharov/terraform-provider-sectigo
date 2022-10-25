package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
func (c *Client) EnrollCertificate(certificate *CertificateRequest) (*CertificateEnrollResponse, error) {

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/ssl/v1/enroll", c.URL),
		"",
		)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := CertificateEnrollResponse{}
	err = json.Unmarshal(body, &ct)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
*/

func (c *Client) GetCertificateTypes() (*[]CertificateTypes, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/v1/types", c.URL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := make([]CertificateTypes, 0)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetCertificateCustomFieldDefinitions() (*[]CertificateCustomFieldDefinition, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/v1/customFields", c.URL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := make([]CertificateCustomFieldDefinition, 0)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
