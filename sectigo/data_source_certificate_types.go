package sectigo

import (
	"context"
	"strconv"
	"terraform-provider-sectigo/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCertificateTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertificateTypesRead,
		Schema: map[string]*schema.Schema{
			"certificate_types": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"use_secondary_org_name": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"terms": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"key_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func flattenCertificateTypes(t *[]sdk.CertificateTypes) []interface{} {
	r := make([]interface{}, len(*t), len(*t))

	for i, types := range *t {
		certTypes := make(map[string]interface{})
		certTypes["id"] = types.Id
		certTypes["name"] = types.Name
		certTypes["description"] = types.Description
		certTypes["terms"] = types.Terms

		keyTypes := make([]interface{}, len(types.KeyTypes), len(types.KeyTypes))

		j := 0
		for k, v := range types.KeyTypes {
			keyType := make(map[string]interface{})
			keyType["type"] = k
			keyType["values"] = v
			keyTypes[j] = keyType
			j++
		}
		certTypes["key_types"] = keyTypes

		r[i] = certTypes
	}

	return r
}

func dataSourceCertificateTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	c := m.(*sdk.Client)
	types, err := c.GetCertificateTypes()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("certificate_types", flattenCertificateTypes(types)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
