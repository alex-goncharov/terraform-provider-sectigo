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
	_ datasource.DataSource              = &certificateTypesDataSource{}
	_ datasource.DataSourceWithConfigure = &certificateTypesDataSource{}
)

func NewCertificateTypesDataSource() datasource.DataSource {
	return &certificateTypesDataSource{}
}

type certificateTypesDataSource struct {
	client *sdk.Client
}

type certificateTypesDataSourceModel struct {
	Items []certificateTypeModel `tfsdk:"items"`
}

type certificateTypeModel struct {
	ID                  types.Int64    `tfsdk:"id"`
	Name                types.String   `tfsdk:"name"`
	Description         types.String   `tfsdk:"description"`
	UseSecondaryOrgName types.Bool     `tfsdk:"use_secondary_org_name"`
	Terms               []types.Int64  `tfsdk:"terms"`
	KeyTypes            []keyTypeModel `tfsdk:"key_types"`
}

type keyTypeModel struct {
	Type  types.String   `tfsdk:"type"`
	Sizes []types.String `tfsdk:"sizes"`
}

func (d *certificateTypesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_types"
}

func (d *certificateTypesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches certificate types",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Description: "List of certificate types",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Certificate type ID",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Certificate type name",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Certificate type description",
							Computed:    true,
						},
						"use_secondary_org_name": schema.BoolAttribute{
							Description: "Should secondary organization name be used",
							Computed:    true,
						},
						"terms": schema.SetAttribute{
							Description: "Terms?",
							Computed:    true,
							ElementType: types.Int64Type,
						},
						"key_types": schema.ListNestedAttribute{
							Description: "Allowed key types",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description: "Key type",
										Computed:    true,
									},
									"sizes": schema.SetAttribute{
										Description: "List of allowed key sizes",
										Computed:    true,
										ElementType: types.StringType,
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

func (d *certificateTypesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*sdk.Client)
}

func (d *certificateTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state certificateTypesDataSourceModel

	t, err := d.client.GetCertificateTypes()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching Sectigo certificate types",
			fmt.Sprintf("Could not read Sectigo certificate types: %v", err),
		)
		return
	}

	for _, _t := range *t {
		ct := certificateTypeModel{
			ID:                  types.Int64Value(_t.Id),
			Name:                types.StringValue(_t.Name),
			Description:         types.StringValue(_t.Description),
			UseSecondaryOrgName: types.BoolValue(_t.UseSecondaryOrgName),
		}
		for _, term := range _t.Terms {
			ct.Terms = append(ct.Terms, types.Int64Value(term))
		}

		for k, v := range _t.KeyTypes {
			kt := keyTypeModel{
				Type: types.StringValue(k),
			}
			for _, _v := range v {
				kt.Sizes = append(kt.Sizes, types.StringValue(_v))

			}
			ct.KeyTypes = append(ct.KeyTypes, kt)
		}

		state.Items = append(state.Items, ct)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
