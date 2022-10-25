package sectigo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-sectigo/sdk"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECTIGO_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECTIGO_PASSWORD", nil),
			},
			"customer_uri": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECTIGO_CUSTOMER_URI", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"sectigo_certificate_types":                    dataSourceCertificateTypes(),
			"sectigo_certificate_custom_field_definitions": dataSourceCertificateCustomFieldDefinitions(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	customerUri := d.Get("customer_uri").(string)

	var diags diag.Diagnostics

	if (username != "") && (password != "") && (customerUri != "") {
		c := sdk.NewClient(&username, &customerUri, &password)
		return c, diags
	}

	return nil, nil // FIXME: do something sensible
}
