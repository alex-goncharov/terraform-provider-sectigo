package sdk

type CertificateTypes struct {
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName"`
	Terms               []int               `json:"terms"`
	KeyTypes            map[string][]string `json:"keyTypes"`
}

type CertificateCustomFieldDefinition struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Mandatory bool   `json:"mandatory"`
}

type CertificateCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CertificateDetails struct {
	Issuer          string `json:"issuer"`
	Subject         string `json:"subject"`
	SubjectAltNames string `json:"subjectAltNames"`
	Md5Hash         string `json:"md5Hash"`
	Sha1Hash        string `json:"sha1Hash"`
}

type Certificate struct {
	CommonName           string                   `json:"commonName"`
	SslId                int                      `json:"sslId"`
	Id                   int                      `json:"id"`
	OrgId                int                      `json:"orgId"`
	Status               string                   `json:"status"`
	OrderNumber          int                      `json:"orderNumber"`
	BackendCertId        string                   `json:"backendCertId"`
	Vendor               string                   `json:"vendor"`
	CertType             CertificateTypes         `json:"certType"`
	SubType              string                   `json:"subType"`
	Term                 int                      `json:"term"`
	Owner                string                   `json:"owner"`
	OwnerId              int                      `json:"ownerId"`
	Requester            string                   `json:"requester"`
	RequesterId          int                      `json:"requesterId"`
	RequestedVia         string                   `json:"requestedVia"`
	Comments             string                   `json:"comments"`
	Requested            string                   `json:"requested"`
	Approved             string                   `json:"approved"`
	Issued               string                   `json:"issued"`
	Expires              string                   `json:"expires"`
	Renewed              bool                     `json:"renewed"`
	SerialNumber         string                   `json:"serialNumber"`
	SignatureAlg         string                   `json:"signatureAlg"`
	KeyAlgorithm         string                   `json:"keyAlgorithm"`
	KeySize              int                      `json:"keySize"`
	KeyType              string                   `json:"keyType"`
	KeyUsages            []string                 `json:"keyUsages"`
	ExtendedKeyUsages    []string                 `json:"extendedKeyUsages"`
	SuspendNotifications bool                     `json:"suspendNotifications"`
	CustomFields         []CertificateCustomField `json:"customFields"`
	CertificateDetails   CertificateDetails

	AutoInstallDetails struct {
		State string `json:"state"`
	} `json:"autoInstallDetails"`

	AutoRenewDetails struct {
		State string `json:"state"`
	} `json:"autoRenewDetails"`
}

type CertificateRequest struct {
	OrgId             int                      `json:"orgId"`
	SubjAltNames      string                   `json:"subjAltNames"`
	CertType          int                      `json:"certType"`
	NumberServers     int                      `json:"numberServers"`
	ServerType        int                      `json:"serverType"`
	Term              int                      `json:"term"`
	Comments          string                   `json:"comments"`
	ExternalRequester string                   `json:"externalRequester"`
	CustomFields      []CertificateCustomField `json:"customFields"`
	Csr               string                   `json:"csr"`
}

type CertificateEnrollResponse struct {
	Id      int    `json:"sslId"`
	RenewId string `json:"renewId"`
}

type CertificateListItem struct {
	Id           int    `json:"sslId"`
	CommonName   string `json:"commonName"`
	SerialNumber string `json:"serialNumber"`
}

type CertificateList []CertificateListItem
