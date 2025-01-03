package provider

import (
	"context"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithFunctions = &ClearBladeGoogleProvider{}

// ClearBladeGoogleProvider defines the provider implementation.
type ClearBladeGoogleProvider struct{}

// ClearBladeGoogleProviderModel describes the provider data model.
type ClearBladeGoogleProviderModel struct {
	Project types.String `tfsdk:"project"`
	Region  types.String `tfsdk:"region"`
}

func (o *ClearBladeGoogleProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "clearblade-google"
}

func (o *ClearBladeGoogleProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project": schema.StringAttribute{
				Optional: true,
			},
			"region": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (o *ClearBladeGoogleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ClearBladeGoogleProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create secret mgr client", err.Error())
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (o *ClearBladeGoogleProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewMEKResource,
		NewRandomStringResource,
	}
}

func (o *ClearBladeGoogleProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewHelmValuesDataSource,
	}
}

func (o *ClearBladeGoogleProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &ClearBladeGoogleProvider{}
	}
}
