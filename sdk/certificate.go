package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) EnrollCertificate(csr *CertificateRequest) (*CertificateEnrollResponse, error) {

	payload, err := json.Marshal(csr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/ssl/v1/enroll", c.URL),
		bytes.NewReader(payload),
	)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := CertificateEnrollResponse{}
	err = json.Unmarshal(body, &r)
	return &r, err
}

func (c *Client) GetCertificateTypes() (*[]CertificateType, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/v1/types", c.URL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := make([]CertificateType, 0)
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

func (c *Client) ListCertificates() (*CertificateList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/v1", c.URL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := make(CertificateList, 0)
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (c *Client) GetCertificate(id int64) (*Certificate, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/v1/%d", c.URL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var r Certificate
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
