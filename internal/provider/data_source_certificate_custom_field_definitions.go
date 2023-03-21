package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-sectigo/sdk"
)

var (
	_ datasource.DataSource              = &certificateCustomFieldDefinitionsDataSource{}
	_ datasource.DataSourceWithConfigure = &certificateCustomFieldDefinitionsDataSource{}
)

func NewCertificateCustomFieldDefinitionsDataSource() datasource.DataSource {
	return &certificateCustomFieldDefinitionsDataSource{}
}

type certificateCustomFieldDefinitionsDataSource struct {
	client *sdk.Client
}

type certificateCustomFieldDefinitionsModel struct {
	Items []certificateCustomFieldDefinitionModel `tfsdk:"items"`
}

type certificateCustomFieldDefinitionModel struct {
	ID        types.Int64  `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Mandatory types.Bool   `tfsdk:"mandatory"`
}

func (d *certificateCustomFieldDefinitionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_custom_field_definitions"
}

func (d *certificateCustomFieldDefinitionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches SSL certificate custom field definitions",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Description: "List of SSL certificate custom field definitions",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Custom field definition ID",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Custom field definition name",
							Computed:    true,
						},
						"mandatory": schema.BoolAttribute{
							Description: "Is the use of the custom field mandatory",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *certificateCustomFieldDefinitionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*sdk.Client)
}

func (d *certificateCustomFieldDefinitionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state certificateCustomFieldDefinitionsModel

	defs, err := d.client.GetCertificateCustomFieldDefinitions()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching Sectigo certificate custom field definitions",
			fmt.Sprintf("Could not read Sectigo certificate custom field definitions: %v", err),
		)
		return
	}

	for _, _d := range *defs {
		def := certificateCustomFieldDefinitionModel{
			ID:        types.Int64Value(_d.Id),
			Name:      types.StringValue(_d.Name),
			Mandatory: types.BoolValue(_d.Mandatory),
		}
		state.Items = append(state.Items, def)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
