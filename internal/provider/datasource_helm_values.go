package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gopkg.in/yaml.v3"
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
				MarkdownDescription: "Helm values",
				Computed:            true,
			},
			"options": schema.SingleNestedAttribute{
				MarkdownDescription: "Infrastructure options",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"global": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart global section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"namespace": schema.StringAttribute{
								MarkdownDescription: "Instance namespace to deploy to",
								Required:            true,
							},
							"image_puller_secret": schema.StringAttribute{
								MarkdownDescription: "Image puller secret key needed to pull the container images from GCR",
								Required:            true,
							},
							"enterprise_base_url": schema.StringAttribute{
								MarkdownDescription: "The base URL the platform will be reachable at",
								Required:            true,
							},
							"enterprise_blue_version": schema.StringAttribute{
								MarkdownDescription: "The blue version is the default ClearBlade version",
								Required:            true,
							},
							"enterprise_green_version": schema.StringAttribute{
								MarkdownDescription: "Utilized during Blue/Green upgrades. Leave blank if not using",
								Required:            false,
								Optional:            true,
							},
							"enterprise_console_version": schema.StringAttribute{
								MarkdownDescription: "The console version is the default Cb Console version",
								Required:            true,
							},
							"enterprise_slot": schema.StringAttribute{
								MarkdownDescription: "Utilized during Blue/Green upgrades. Leave blank if not using",
								Required:            false,
								Optional:            true,
							},
							"enterprise_instance_id": schema.StringAttribute{
								MarkdownDescription: "The Instance ID for the deployment, provided by ClearBlade",
								Required:            true,
							},
							"enterprise_registration_key": schema.StringAttribute{
								MarkdownDescription: "Unique registration key for new users to register with the platform",
								Required:            true,
							},
							"gcp_project": schema.StringAttribute{
								MarkdownDescription: "GCP project ID",
								Required:            true,
							},
							"gcp_region": schema.StringAttribute{
								MarkdownDescription: "GCP region",
								Required:            true,
							},
							"gcp_gsm_service_account": schema.StringAttribute{
								MarkdownDescription: "Google Secret Manager service account email",
								Required:            true,
							},
							"storage_class_name": schema.StringAttribute{
								MarkdownDescription: "The storage class used by all Persistent Volume Claims in the deployment",
								Required:            true,
							},
							"iotcore_enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to true if this deployment uses the IOTCore Sidecar",
								Required:            true,
							},
							"ia_enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to true if this deployment uses the Intelligent Assets Sidecar",
								Required:            true,
							},
							"ops_console_enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to true if this deployment uses the Ops Console Sidecar",
								Required:            true,
							},
							"gcp_cloudsql_enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to true if you are using GCP's Cloud SQL instead of postgres",
								Required:            true,
							},
							"gcp_memorystore_enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to true if you are using GCP's MemoryStore instead of redis",
								Required:            true,
							},
						},
					},
					"cb_console": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb-console section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
						},
					},
					"cb_file_hosting": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_file_hosting section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
						},
					},
					"cb_haproxy": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_haproxy section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"replicas": schema.Int32Attribute{
								MarkdownDescription: "Number of HAProxy replicas",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to false if using an external HAProxy deployment",
								Required:            true,
							},
							"primary_ip": schema.StringAttribute{
								MarkdownDescription: "Required if using this HAProxy. The primary external IP address for the deployment",
								Required:            true,
							},
							"mqtt_ip": schema.StringAttribute{
								MarkdownDescription: "Required if utilizing external MQTT connections",
								Required:            true,
							},
							"mqtt_over_443": schema.BoolAttribute{
								MarkdownDescription: "Set to true if you would like MQTT connections to work over 443 in addition to the default 1883",
								Required:            true,
							},
							"cert_renewal": schema.BoolAttribute{
								MarkdownDescription: "Set to true for automatic cert renewal with LetsEncrypt",
								Required:            true,
							},
							"renewal_days": schema.Int32Attribute{
								MarkdownDescription: "Days out to renew cert",
								Required:            true,
							},
							"controller_version": schema.StringAttribute{
								MarkdownDescription: "Image tag of the cb controller",
								Required:            true,
							},
							"acme_config": schema.ListNestedAttribute{
								MarkdownDescription: "ACME config",
								Required:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"directory": schema.StringAttribute{
											MarkdownDescription: "Directory of the ACME config",
											Required:            true,
										},
										"email": schema.StringAttribute{
											MarkdownDescription: "Email of the ACME config",
											Required:            true,
										},
										"eab_kid": schema.StringAttribute{
											MarkdownDescription: "EAB KID of the ACME config",
											Required:            true,
										},
										"eab_key": schema.StringAttribute{
											MarkdownDescription: "EAB Key of the ACME config",
											Required:            true,
										},
										"key_type": schema.StringAttribute{
											MarkdownDescription: "Key type of the ACME config",
											Required:            true,
										},
										"domain": schema.StringAttribute{
											MarkdownDescription: "Domain to request certificate for",
											Required:            true,
										},
									},
								},
							},
						},
					},
					"cb_iotcore": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_iotcore section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"check_clearblade_readiness": schema.BoolAttribute{
								MarkdownDescription: "Set to true to force the IOTCore pod to wait for the Clearblade pods before starting",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
							"version": schema.StringAttribute{
								MarkdownDescription: "Version of the IOTCore",
								Required:            false,
								Optional:            true,
							},
							"regions": schema.StringAttribute{
								MarkdownDescription: "Regions to deploy the IOTCore to",
								Required:            false,
								Optional:            true,
							},
						},
					},
					"cb_ia": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_ia section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"check_clearblade_readiness": schema.BoolAttribute{
								MarkdownDescription: "Set to true to force the Intelligent Assets pod to wait for the Clearblade pods before starting",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
							"version": schema.StringAttribute{
								MarkdownDescription: "Version of Intelligent Assets",
								Required:            false,
								Optional:            true,
							},
						},
					},
					"cb_postgres": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_postgres section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to false if using an external postgres deployment",
								Required:            true,
							},
							"replicas": schema.Int32Attribute{
								MarkdownDescription: "Number of Postgres replicas",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
							"postgres0_disk_name": schema.StringAttribute{
								MarkdownDescription: "Postgres0 disk name",
								Required:            true,
							},
						},
					},
					"cb_redis": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart cb_redis section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								MarkdownDescription: "Set to false if using an external redis deployment",
								Required:            true,
							},
							"high_availability": schema.BoolAttribute{
								MarkdownDescription: "Set to true to utilize redis sentinel with automatic failover. Requires roughly 4x CPU/mem as a non-HA deployment",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
						},
					},
					"clearblade": schema.SingleNestedAttribute{
						MarkdownDescription: "Helm Chart clearblade section",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"blue_replicas": schema.Int32Attribute{
								MarkdownDescription: "If not using blue/green deployments, blue is the default",
								Required:            true,
							},
							"green_replicas": schema.Int32Attribute{
								MarkdownDescription: "If not using blue/green deployments, set to 0",
								Required:            true,
							},
							"mqtt_allow_duplicate_client_id": schema.BoolAttribute{
								MarkdownDescription: "Set to true to allow duplicate client IDs. Set to false to reject duplicate connections",
								Required:            true,
							},
							"request_cpu": schema.Float32Attribute{
								MarkdownDescription: "Requested CPUs",
								Required:            true,
							},
							"request_memory": schema.StringAttribute{
								MarkdownDescription: "Requested memory",
								Required:            true,
							},
							"limit_cpu": schema.Float32Attribute{
								MarkdownDescription: "CPU limit",
								Required:            true,
							},
							"limit_memory": schema.StringAttribute{
								MarkdownDescription: "Memory limit",
								Required:            true,
							},
							"license_renewal_webhooks": schema.ListAttribute{
								MarkdownDescription: "List of webhooks to request for license renewal",
								Required:            true,
								ElementType:         types.StringType,
							},
							"metrics_reporting_webhooks": schema.ListAttribute{
								MarkdownDescription: "List of webhooks to attempt reporting metrics to",
								Required:            true,
								ElementType:         types.StringType,
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
	values, diags := data.Options.toHelmValues()
	if len(diags) > 0 {
		resp.Diagnostics.Append(diags...)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	valuesYaml, err := yaml.Marshal(values)
	if err != nil {
		resp.Diagnostics.AddError("Failed to marhsal values", err.Error())
		return
	}
	data.HelmValues = types.StringValue(string(valuesYaml))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
