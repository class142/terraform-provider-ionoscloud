package s3management

import (
	"context"
	"fmt"

	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services/s3management"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigure = (*accessKeyDataSource)(nil)

// NewBucketDataSource creates a new data source for the accesskey resource.
func NewBucketDataSource() datasource.DataSource {
	return &accessKeyDataSource{}
}

type accessKeyDataSource struct {
	client *services.SdkBundle
}

// Metadata returns the metadata for the data source.
func (d *accessKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_s3_accesskey"
}

// Configure configures the data source.
func (d *accessKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*services.SdkBundle)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *s3.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

// Schema returns the schema for the data source.
func (d *accessKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:    true,
				Description: "The ID (UUID) of the AccessKey.",
			},
			"description": schema.StringAttribute{
				Description: "Description of the Access key.",
				Optional:    true,
			},
			"accesskey": schema.StringAttribute{
				Description: "Access key metadata is a string of 92 characters.",
				Computed:    true,
			},
			"canonical_user_id": schema.StringAttribute{
				Description: "The canonical user ID which is valid for user-owned buckets.",
				Computed:    true,
			},
			"contract_user_id": schema.StringAttribute{
				Description: "The contract user ID which is valid for contract-owned buckets",
				Computed:    true,
			},
		},
	}
}

// Read reads the data source.
func (d *accessKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	if d.client == nil {
		resp.Diagnostics.AddError("s3 api client not configured", "The provider client is not configured")
		return
	}

	var data *s3management.AccessKeyDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// result, found, err := d.client.GetAccessKeyForDataSource(ctx, data.Name)
	// if err != nil {
	// 	resp.Diagnostics.AddError("failed to get accesskey", err.Error())
	// 	return
	// }

	if !found {
		resp.Diagnostics.AddError("accesskey not found", "The accesskey was not found")
		return
	}

	data = result
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
