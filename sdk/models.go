package sdk

type CertificateTypes struct {
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName"`
	Terms               []int               `json:"terms"`
	KeyTypes            map[string][]string `json:"keyTypes"`
}

type CustomFieldDefinition struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Mandatory bool   `json:"mandatory"`
}

type CustomFieldDefinitions []CustomFieldDefinition

type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CertificateRequest struct {
	OrgId             int
	CSR               string
	SubjAltNames      []string
	CertType          int
	Term              int
	Comments          string
	CustomFields      []CustomField
	ExternalRequested string
}

type certificateRequestBody struct {
	OrgId             int           `json:"orgId"`
	SubjAltNames      string        `json:"subjAltNames"`
	CertType          int           `json:"certType"`
	NumberServers     int           `json:"numberServers"`
	ServerType        int           `json:"serverType"`
	Term              int           `json:"term"`
	Comments          string        `json:"comments"`
	ExternalRequester string        `json:"externalRequester"`
	CustomFields      []CustomField `json:"customFields"`
	Csr               string        `json:"csr"`
}

type CertificateEnrollResponse struct {
	ID      string `json:"sslId"`
	RenewID string `json:"renewId"`
}
