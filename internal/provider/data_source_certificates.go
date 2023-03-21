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
	_ datasource.DataSource              = &certificatesDataSource{}
	_ datasource.DataSourceWithConfigure = &certificatesDataSource{}
)

func NewCertificatesDataSource() datasource.DataSource {
	return &certificatesDataSource{}
}

type certificatesDataSource struct {
	client *sdk.Client
}

type certificatesDataSourceModel struct {
	Items []certificateModel `tfsdk:"items"`
}

type certificateModel struct {
	ID           types.Int64  `tfsdk:"id"`
	CommonName   types.String `tfsdk:"common_name"`
	SerialNumber types.String `tfsdk:"serial_number"`
}

func (d *certificatesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificates"
}

func (d *certificatesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches certificates",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Description: "List of certificates",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Certificate type ID",
							Computed:    true,
						},
						"common_name": schema.StringAttribute{
							Description: "Certificate type name",
							Computed:    true,
						},
						"serial_number": schema.StringAttribute{
							Description: "Certificate type name",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *certificatesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*sdk.Client)
}

func (d *certificatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state certificatesDataSourceModel

	t, err := d.client.ListCertificates()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching Sectigo certificates",
			fmt.Sprintf("Could not read Sectigo certificates: %v", err),
		)
		return
	}

	for _, _t := range *t {
		ct := certificateModel{
			ID:           types.Int64Value(_t.Id),
			CommonName:   types.StringValue(_t.CommonName),
			SerialNumber: types.StringValue(_t.SerialNumber),
		}
		state.Items = append(state.Items, ct)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
