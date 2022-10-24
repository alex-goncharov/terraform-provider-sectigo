package sectigo

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-sectigo/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type certificateTypesResponse struct {
	Id                  int                 `json:"id"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	UseSecondaryOrgName bool                `json:"useSecondaryOrgName"`
	Terms               []int               `json:"terms"`
	KeyTypes            map[string][]string `json:"keyTypes"`
}

type certificateKeyTypes struct {
	Type   string
	Values []string
}

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

func flattenCertificateTypes(t *sdk.CertificateTypes) {

}

func dataSourceCertificateTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	c := m.(*sdk.Client)
	types, err := c.GetCertificateTypes()

	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("url %v", c.URL, c))
		return diag.FromErr(err)
	}

	/*
			sdk.NewClient("dataworks_acmeuser_nonprod_eai_3536187",
		 "fedex-dev",
		 "quuay0eing7MieR3geisueNgac5choh!"
		)
		req.Header.Set("login", "dataworks_acmeuser_nonprod_eai_3536187")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("customerUri", "fedex-dev")
		req.Header.Set("password", "quuay0eing7MieR3geisueNgac5choh!")
	*/

	xx := make([]interface{}, len(*types), len(*types))

	for i, x := range *types {
		xxx := make(map[string]interface{})
		xxx["id"] = x.Id
		xxx["name"] = x.Name
		xxx["description"] = x.Description
		xxx["terms"] = x.Terms

		keyTypes := make([]interface{}, len(x.KeyTypes), len(x.KeyTypes))

		j := 0
		for k, v := range x.KeyTypes {
			zz := make(map[string]interface{})
			zz["type"] = k
			zz["values"] = v
			keyTypes[j] = zz
			j++
		}
		xxx["key_types"] = keyTypes

		xx[i] = xxx
		tflog.Debug(ctx, fmt.Sprintf("XXX value %v", xxx))
		tflog.Debug(ctx, fmt.Sprintf("XXX value %v", xx))
	}

	if err := d.Set("certificate_types", xx); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
