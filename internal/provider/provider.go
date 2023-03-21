package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	"terraform-provider-sectigo/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &SectigoProvider{}

type SectigoProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SectigoProviderModel describes the provider data model.
type SectigoProviderModel struct {
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
	CustomerUri types.String `tfsdk:"customer_uri"`
}

func (p *SectigoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sectigo"
	resp.Version = p.version
}

func (p *SectigoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description:         "Sectigo API user name.",
				MarkdownDescription: "Sectigo API user name.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				Description:         "Sectigo API password.",
				MarkdownDescription: "Sectigo API password.",
				Required:            true,
				Sensitive:           true,
			},
			"customer_uri": schema.StringAttribute{
				Description:         "Sectigo API customer URI.",
				MarkdownDescription: "Sectigo API customer URI.",
				Required:            true,
			},
		},
	}
}

func (p *SectigoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SectigoProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Sectigo API Username",
			"The provider cannot create the Sectigo API client as there is an unknown configuration value for the Sectigo API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SECTIGO_USERNAME environment variable.",
		)
	}

	if config.CustomerUri.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("customer_uri"),
			"Unknown Sectigo API CustomerUri",
			"The provider cannot create the Sectigo API client as there is an unknown configuration value for the Sectigo API customer URI. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SECTIGO_CUSTOMER_URI environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Sectigo API Password",
			"The provider cannot create the Sectigo API client as there is an unknown configuration value for the Sectigo API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SECTIGO_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	username := os.Getenv("SECTIGO_USERNAME")
	customerUri := os.Getenv("SECTIGO_CUSTOMER_UR")
	password := os.Getenv("SECTIGO_PASSWORD")

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}
	if !config.CustomerUri.IsNull() {
		customerUri = config.CustomerUri.ValueString()
	}
	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Sectigo API Username",
			"The provider cannot create the Sectigo API client as there is a missing or empty value for the Sectigo API username. "+
				"Set the username value in the configuration or use the SECTIGO_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if customerUri == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("customer_uri"),
			"Missing Sectigo API customer URI",
			"The provider cannot create the Sectigo API client as there is a missing or empty value for the Sectigo API customer URI. "+
				"Set the customer_uri value in the configuration or use the SECTIGO_CUSTOMER_URI environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Sectigo API password",
			"The provider cannot create the Sectigo API client as there is a missing or empty value for the Sectigo API password. "+
				"Set the password value in the configuration or use the SECTIGO_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "sectigo_username", username)
	ctx = tflog.SetField(ctx, "sectigo_customer_uri", customerUri)
	ctx = tflog.SetField(ctx, "sectigo_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "sectigo_password", password)

	tflog.Debug(ctx, "Creating Sectigo API client")
	c := sdk.NewClient(&username, &customerUri, &password)

	tflog.Error(ctx, fmt.Sprintf("AWW NO ERRORS client is %v", c))

	resp.DataSourceData = c
	resp.ResourceData = c

	tflog.Info(ctx, "Configured Sectigo API client", map[string]any{"success": true})
}

func (p *SectigoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *SectigoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCertificateDataSource,
		NewCertificatesDataSource,
		NewCertificateTypesDataSource,
		NewCertificateCustomFieldDefinitionsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SectigoProvider{
			version: version,
		}
	}
}
