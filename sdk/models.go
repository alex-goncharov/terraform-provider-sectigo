package sdk

type CertificateType struct {
	Id                  int64               `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName"`
	Terms               []int64             `json:"terms"`
	KeyTypes            map[string][]string `json:"keyTypes"`
}

type CertificateCustomFieldDefinition struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Mandatory bool   `json:"mandatory"`
}

const (
	FormatX509    = "x509"    // Certificate (w/ chain), PEM encoded
	FormatX509CO  = "x509CO"  // Certificate only, PEM encoded
	FormatBase64  = "base64"  // PKCS#7, PEM encoded
	FormatBin     = "bin"     // PKCS#7, 'x509IO' - for Root/Intermediate(s) only, PEM encoded
	FormatX509IOR = "x509IOR" // Intermediate(s)/Root only, PEM encoded
	FormatPem     = "pem"     // Certificate (w/ chain), PEM encoded
	FormatPemCO   = "pemco"   // Certificate only, PEM encoded
)

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
	SslId                int64                    `json:"sslId"`
	Id                   int64                    `json:"id"`
	OrgId                int64                    `json:"orgId"`
	Status               string                   `json:"status"`
	OrderNumber          int64                    `json:"orderNumber"`
	BackendCertId        string                   `json:"backendCertId"`
	Vendor               string                   `json:"vendor"`
	CertType             CertificateType          `json:"certType"`
	SubType              string                   `json:"subType"`
	Term                 int64                    `json:"term"`
	Owner                string                   `json:"owner"`
	OwnerId              int64                    `json:"ownerId"`
	Requester            string                   `json:"requester"`
	RequesterId          int64                    `json:"requesterId"`
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
	KeySize              int64                    `json:"keySize"`
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
	OrgId             int64                    `json:"orgId"`
	SubjAltNames      string                   `json:"subjAltNames"`
	CertType          int64                    `json:"certType"`
	NumberServers     int64                    `json:"numberServers"`
	ServerType        int64                    `json:"serverType"`
	Term              int64                    `json:"term"`
	Comments          string                   `json:"comments"`
	ExternalRequester string                   `json:"externalRequester"`
	CustomFields      []CertificateCustomField `json:"customFields"`
	Csr               string                   `json:"csr"`
}

type CertificateEnrollResponse struct {
	Id      int64  `json:"sslId"`
	RenewId string `json:"renewId"`
}

type CertificateListItem struct {
	Id           int64  `json:"sslId"`
	CommonName   string `json:"commonName"`
	SerialNumber string `json:"serialNumber"`
}

type CertificateList []CertificateListItem
