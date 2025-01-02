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

var _ provider.ProviderWithFunctions = &OneClickDeployProvider{}

// OneClickDeployProvider defines the provider implementation.
type OneClickDeployProvider struct{}

// OneClickDeployProviderModel describes the provider data model.
type OneClickDeployProviderModel struct {
	Project types.String `tfsdk:"project"`
	Region  types.String `tfsdk:"region"`
}

func (o *OneClickDeployProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "clearblade-google"
}

func (o *OneClickDeployProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
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

func (o *OneClickDeployProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data OneClickDeployProviderModel

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

func (o *OneClickDeployProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewMEKResource,
		NewRandomStringResource,
	}
}

func (o *OneClickDeployProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (o *OneClickDeployProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewGetHelmYamlTemplateFunction,
	}
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &OneClickDeployProvider{}
	}
}
