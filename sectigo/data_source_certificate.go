package sectigo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-sectigo/sdk"
)

func dataSourceCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertificateRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenCertificateDetails(t sdk.CertificateDetails) []interface{} {
	r := make([]interface{}, 1, 1)
	field := make(map[string]interface{})
	field["issuer"] = t.Issuer
	r[0] = field
	return r
}

func dataSourceCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	c := m.(*sdk.Client)
	r, err := c.GetCertificate(d.Get("id").(int))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(r.Id))
	d.Set("status", r.Status)
	d.Set("serial_number", r.SerialNumber)
	d.Set("certificate_details", flattenCertificateDetails(r.CertificateDetails))

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
