package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &HelmValuesDataSource{}

func NewHelmValuesDataSource() datasource.DataSource {
	return &HelmValuesDataSource{}
}

// HelmValuesDataSource defines the data source implementation.
type HelmValuesDataSource struct{}

// HelmValuesDataSourceModel describes the data source data model.
type HelmValuesDataSourceModel struct {
	Options    TfHelmValues `tfsdk:"options"`
	HelmValues types.String `tfsdk:"values"`
}

func (d *HelmValuesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_helm_values"
}

func (d *HelmValuesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Datasource for creating helm values",

		Attributes: map[string]schema.Attribute{
			"values": schema.StringAttribute{
				MarkdownDescription: "Final Helm values",
				Computed:            true,
			},
			"options": schema.SingleNestedAttribute{
				MarkdownDescription: "Helm values",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"global": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart global section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"imagePullerSecret": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"enterpriseBaseURL": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"enterpriseBlueVersion": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"enterpriseInstanceID": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"enterpriseRegistrationKey": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"gcpProject": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"gcpRegion": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"gcpGSMServiceAccount": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"storageClassName": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"iotCoreEnabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"iaEnabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"gcpCloudSQLEnabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"gcpMemoryStoreEnabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-console": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-console section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-file-hosting": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-file-hosting section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-haproxy": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-haproxy section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"replicas": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"primaryIP": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"mqttIP": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"mqttOver443": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-iotcore": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-iotcore section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"checkClearbladeReadiness": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-ia": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-ia section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"checkClearbladeReadiness": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-postgres": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-postgres section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"replicas": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"postgres0DiskName": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"cb-redis": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-redis section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"highAvailability": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
					"clearblade": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart clearblade section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"blueReplicas": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"greenReplicas": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"mqttAllowDuplicateClientID": schema.BoolAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"requestMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitCPU": schema.Int32Attribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
							"limitMemory": schema.StringAttribute{
								MarkdownDescription: "Instance namespace",
								Required:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *HelmValuesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
}

func (d *HelmValuesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HelmValuesDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	data.Options.Global.Cloud = types.StringValue("gcp")
	valuesStr, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("Failed to marhsal data", err.Error())
		return
	}
	data.HelmValues = types.StringValue(string(valuesStr))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
