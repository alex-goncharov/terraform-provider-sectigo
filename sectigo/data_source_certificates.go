package sectigo

import (
	"context"
	"strconv"
	"terraform-provider-sectigo/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCertificates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertificatesRead,
		Schema: map[string]*schema.Schema{
			"certificates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"common_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenCertificates(t *sdk.CertificateList) []interface{} {
	r := make([]interface{}, len(*t), len(*t))

	for i, f := range *t {
		field := make(map[string]interface{})
		field["id"] = f.Id
		field["common_name"] = f.CommonName
		field["serial_number"] = f.SerialNumber
		r[i] = field
	}

	return r
}

func dataSourceCertificatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	c := m.(*sdk.Client)
	r, err := c.ListCertificates()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("certificates", flattenCertificates(r)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
