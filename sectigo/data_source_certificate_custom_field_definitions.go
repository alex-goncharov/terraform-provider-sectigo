package sectigo

import (
	"context"
	"strconv"
	"terraform-provider-sectigo/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCertificateCustomFieldDefinitions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertificateCustomFieldDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"definitions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mandatory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenCertificateCustomFieldDefinitions(t *[]sdk.CertificateCustomFieldDefinition) []interface{} {
	r := make([]interface{}, len(*t), len(*t))

	for i, f := range *t {
		field := make(map[string]interface{})
		field["id"] = f.Id
		field["name"] = f.Name
		field["mandatory"] = f.Mandatory
		r[i] = field
	}

	return r
}

func dataSourceCertificateCustomFieldDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	c := m.(*sdk.Client)
	r, err := c.GetCertificateCustomFieldDefinitions()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("custom_field_definitions", flattenCertificateCustomFieldDefinitions(r)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
