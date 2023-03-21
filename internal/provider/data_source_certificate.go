package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-sectigo/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &certificateDataSource{}
	_ datasource.DataSourceWithConfigure = &certificateDataSource{}
)

func NewCertificateDataSource() datasource.DataSource {
	return &certificateDataSource{}
}

type certificateDetailsModel struct {
	Issuer types.String `tfsdk:"issuer"`
}

type certificateDataSourceModel struct {
	Id                 types.Int64             `tfsdk:"id"`
	Status             types.String            `tfsdk:"status"`
	SerialNumber       types.String            `tfsdk:"serial_number"`
	CertificateDetails certificateDetailsModel `tfsdk:"certificate_details"`
}

// coffeesDataSource is the data source implementation.
type certificateDataSource struct {
	client *sdk.Client
}

func (d *certificateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate"
}

func (d *certificateDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches SSL certificate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Certificate ID.",
				Required:    true,
				Computed:    false,
			},

			"status": schema.StringAttribute{
				Description: "Certificate status.",
				Computed:    true,
			},

			"serial_number": schema.StringAttribute{
				Description: "Certificate serial number.",
				Computed:    true,
			},
			"certificate_details": schema.SingleNestedAttribute{
				Description: "Certificate details.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"issuer": schema.StringAttribute{
						Description: "Certificate issuer.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *certificateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdk.Client)
}

func (d *certificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state certificateDataSourceModel
	var id types.Int64

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cert, err := d.client.GetCertificate(id.ValueInt64())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Sectigo Certificate",
			fmt.Sprintf("Could not read Sectigo certificate Id %d: %v", id, err),
		)
		return
	}

	state.Id = id
	state.Status = types.StringValue(cert.Status)
	state.SerialNumber = types.StringValue(cert.SerialNumber)
	state.CertificateDetails = certificateDetailsModel{
		Issuer: types.StringValue(cert.CertificateDetails.Issuer),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
