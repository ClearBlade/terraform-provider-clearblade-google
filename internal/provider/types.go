package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type HelmValues struct {
	Global        Global        `yaml:"global"`
	Clearblade    Clearblade    `yaml:"clearblade"`
	CbHaproxy     CbHaproxy     `yaml:"cb-haproxy"`
	CbPostgres    CbPostgres    `yaml:"cb-postgres"`
	CbRedis       CbRedis       `yaml:"cb-redis"`
	CbConsole     CbConsole     `yaml:"cb-console"`
	CbFileHosting CbFileHosting `yaml:"cb-file-hosting"`
	CbIotcore     CbIotcore     `yaml:"cb-iotcore"`
}

type Global struct {
	Namespace         string      `yaml:"namespace"`
	NodeSelector      string      `yaml:"nodeSelector"`
	Tolerations       interface{} `yaml:"tolerations"`
	ImagePullerSecret string      `yaml:"imagePullerSecret"`
	Enterprise        Enterprise  `yaml:"enterprise"`
	IotCore           IotCore     `yaml:"iotCore"`
	IA                IA          `yaml:"IA"`
	Gcp               Gcp         `yaml:"gcp"`
	Advanced          Advanced    `yaml:"advanced"`
	MtlsHAProxy       bool        `yaml:"mtlsHAProxy"`
	MtlsClearBlade    bool        `yaml:"mtlsClearBlade"`
	GMP               bool        `yaml:"GMP"`
}

type Enterprise struct {
	Version         string      `yaml:"version"`
	GreenVersion    interface{} `yaml:"greenVersion"`
	Slot            string      `yaml:"slot"`
	BaseURL         string      `yaml:"baseURL"`
	ConsoleURL      interface{} `yaml:"consoleURL"`
	RegistrationKey string      `yaml:"registrationKey"`
	TagOverride     bool        `yaml:"tagOverride"`
	InstanceID      interface{} `yaml:"instanceID"`
}

type IotCore struct {
	Enabled bool   `yaml:"enabled"`
	Version string `yaml:"version"`
	Regions string `yaml:"regions"`
}

type IA struct {
	Enabled bool   `yaml:"enabled"`
	Version string `yaml:"version"`
}

type Gcp struct {
	Project               string `yaml:"project"`
	Region                string `yaml:"region"`
	GsmReadServiceAccount string `yaml:"gsmReadServiceAccount"`
}

type Advanced struct {
	PredefinedNamespace bool        `yaml:"predefinedNamespace"`
	MemoryStore         MemoryStore `yaml:"memoryStore"`
	CloudSQL            CloudSQL    `yaml:"cloudSQL"`
	Secrets             Secrets     `yaml:"secrets"`
}

type MemoryStore struct {
	Enabled bool        `yaml:"enabled"`
	Address interface{} `yaml:"address"`
}

type CloudSQL struct {
	Enabled                bool        `yaml:"enabled"`
	DatabaseConnectionName interface{} `yaml:"databaseConnectionName"`
}

type Secrets struct {
	Manager string `yaml:"manager"`
}

type Clearblade struct {
	Replicas         int              `yaml:"replicas"`
	GreenReplicas    int              `yaml:"greenReplicas"`
	License          License          `yaml:"license"`
	ResourceRequests ResourceRequests `yaml:"resourceRequests"`
	ResourceLimits   ResourceLimits   `yaml:"resourceLimits"`
	Mqtt             Mqtt             `yaml:"mqtt"`
}

type License struct {
	Key        string    `yaml:"key"`
	InstanceID string    `yaml:"instanceID"`
	AutoRenew  AutoRenew `yaml:"autoRenew"`
}

type AutoRenew struct {
	Enabled bool `yaml:"enabled"`
}

type ResourceRequests struct {
	CPU    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type ResourceLimits struct {
	CPU    int    `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type Mqtt struct {
	AllowDuplicateClientID bool `yaml:"allowDuplicateClientID"`
}

type CbHaproxy struct {
	Enabled           bool             `yaml:"enabled"`
	MonitoringEnabled bool             `yaml:"monitoringEnabled"`
	Replicas          int              `yaml:"replicas"`
	Image             string           `yaml:"image"`
	ImageTag          string           `yaml:"imageTag"`
	MqttOver443       bool             `yaml:"mqttOver443"`
	StatsAuth         string           `yaml:"stats_auth"`
	IP                IP               `yaml:"ip"`
	ResourceRequests  ResourceRequests `yaml:"resourceRequests"`
	ResourceLimits    ResourceLimits   `yaml:"resourceLimits"`
}

type IP struct {
	Primary string `yaml:"primary"`
	Mqtt    string `yaml:"mqtt"`
}

type CbPostgres struct {
	Enabled           bool             `yaml:"enabled"`
	MonitoringEnabled bool             `yaml:"monitoringEnabled"`
	Image             string           `yaml:"image"`
	ImageTag          string           `yaml:"imageTag"`
	Replicas          int              `yaml:"replicas"`
	ResourceRequests  ResourceRequests `yaml:"resourceRequests"`
	ResourceLimits    ResourceLimits   `yaml:"resourceLimits"`
}

type CbRedis struct {
	Enabled           bool             `yaml:"enabled"`
	MonitoringEnabled bool             `yaml:"monitoringEnabled"`
	Image             string           `yaml:"image"`
	ImageTag          string           `yaml:"imageTag"`
	ResourceLimits    ResourceLimits   `yaml:"resourceLimits"`
	ResourceRequests  ResourceRequests `yaml:"resourceRequests"`
}

type CbConsole struct {
	ResourceRequests ResourceRequests `yaml:"resourceRequests"`
	ResourceLimits   ResourceLimits   `yaml:"resourceLimits"`
}

type CbFileHosting struct {
	ResourceLimits   ResourceLimits   `yaml:"resourceLimits"`
	ResourceRequests ResourceRequests `yaml:"resourceRequests"`
}

type CbIotcore struct {
	CheckClearbladeReadiness bool           `yaml:"checkClearbladeReadiness"`
	ResourceLimits           ResourceLimits `yaml:"resourceLimits"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Tf

type TfHelmValues struct {
	Global        TfGlobal      `tfsdk:"global"`
	CbConsole     TfConsole     `tfsdk:"cb_console"`
	CbFileHosting TfFileHosting `tfsdk:"cb_file_hosting"`
	CbHaproxy     TfHAProxy     `tfsdk:"cb_haproxy"`
	CbIotcore     TfIotCore     `tfsdk:"cb_iotcore"`
	CbIa          TfIa          `tfsdk:"cb_ia"`
	CbPostgres    TfPostgres    `tfsdk:"cb_postgres"`
	CbRedis       TfRedis       `tfsdk:"cb_redis"`
	Clearblade    TfClearblade  `tfsdk:"clearblade"`
}

type TfGlobal struct {
	Namespace                 types.String `tfsdk:"namespace"`
	ImagePullerSecret         types.String `tfsdk:"image_puller_secret"`
	EnterpriseBaseURL         types.String `tfsdk:"enterprise_base_url"`
	EnterpriseBlueVersion     types.String `tfsdk:"enterprise_blue_version"`
	EnterpriseInstanceID      types.String `tfsdk:"enterprise_instance_id"`
	EnterpriseRegistrationKey types.String `tfsdk:"enterprise_registration_key"`
	IotCoreEnabled            types.Bool   `tfsdk:"iotcore_enabled"`
	IaEnabled                 types.Bool   `tfsdk:"ia_enabled"`
	GcpCloudSQLEnabled        types.Bool   `tfsdk:"gcp_cloudsql_enabled"`
	GcpMemoryStoreEnabled     types.Bool   `tfsdk:"gcp_memorystore_enabled"`
	GcpProject                types.String `tfsdk:"gcp_project"`
	GcpRegion                 types.String `tfsdk:"gcp_region"`
	GcpGSMServiceAccount      types.String `tfsdk:"gcp_gsm_service_account"`
	StorageClassName          types.String `tfsdk:"storage_class_name"`
}

type TfConsole struct {
	RequestCPU    types.Int32  `tfsdk:"request_cpu"`
	RequestMemory types.String `tfsdk:"request_memory"`
	LimitCPU      types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory   types.String `tfsdk:"limit_memory"`
}

type TfFileHosting struct {
	RequestCPU    types.Int32  `tfsdk:"request_cpu"`
	RequestMemory types.String `tfsdk:"request_memory"`
	LimitCPU      types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory   types.String `tfsdk:"limit_memory"`
}

type TfHAProxy struct {
	Replicas      types.Int32  `tfsdk:"replicas"`
	RequestCPU    types.Int32  `tfsdk:"request_cpu"`
	RequestMemory types.String `tfsdk:"request_memory"`
	LimitCPU      types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory   types.String `tfsdk:"limit_memory"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	PrimaryIP     types.String `tfsdk:"primary_ip"`
	MqttIP        types.String `tfsdk:"mqtt_ip"`
	MqttOver443   types.Bool   `tfsdk:"mqtt_over_443"`
}

type TfIotCore struct {
	CheckClearbladeReadiness types.Bool   `tfsdk:"check_clearblade_readiness"`
	RequestCPU               types.Int32  `tfsdk:"request_cpu"`
	RequestMemory            types.String `tfsdk:"request_memory"`
	LimitCPU                 types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory              types.String `tfsdk:"limit_memory"`
}

type TfIa struct {
	CheckClearbladeReadiness types.Bool   `tfsdk:"check_clearblade_readiness"`
	RequestCPU               types.Int32  `tfsdk:"request_cpu"`
	RequestMemory            types.String `tfsdk:"request_memory"`
	LimitCPU                 types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory              types.String `tfsdk:"limit_memory"`
}

type TfPostgres struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	Replicas          types.Int32  `tfsdk:"replicas"`
	RequestCPU        types.Int32  `tfsdk:"request_cpu"`
	RequestMemory     types.String `tfsdk:"request_memory"`
	LimitCPU          types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory       types.String `tfsdk:"limit_memory"`
	Postgres0DiskName types.String `tfsdk:"postgres0_disk_name"`
}

type TfRedis struct {
	Enabled          types.Bool   `tfsdk:"enabled"`
	HighAvailability types.Bool   `tfsdk:"high_availability"`
	RequestCPU       types.Int32  `tfsdk:"request_cpu"`
	RequestMemory    types.String `tfsdk:"request_memory"`
	LimitCPU         types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory      types.String `tfsdk:"limit_memory"`
}

type TfClearblade struct {
	BlueReplicas               types.Int32  `tfsdk:"blue_replicas"`
	GreenReplicas              types.Int32  `tfsdk:"green_replicas"`
	MqttAllowDuplicateClientID types.Bool   `tfsdk:"mqtt_allow_duplicate_client_id"`
	RequestCPU                 types.Int32  `tfsdk:"request_cpu"`
	RequestMemory              types.String `tfsdk:"request_memory"`
	LimitCPU                   types.Int32  `tfsdk:"limit_cpu"`
	LimitMemory                types.String `tfsdk:"limit_memory"`
}

func (t *TfHelmValues) toHelmValues() *HelmValues {
	h := &HelmValues{
		Global: Global{
			Namespace:         t.Global.Namespace.ValueString(),
			ImagePullerSecret: t.Global.ImagePullerSecret.ValueString(),
			Enterprise: Enterprise{
				Version:         t.Global.EnterpriseBlueVersion.ValueString(),
				Slot:            "blue",
				BaseURL:         t.Global.EnterpriseBaseURL.ValueString(),
				RegistrationKey: t.Global.EnterpriseRegistrationKey.ValueString(),
				TagOverride:     false,
			},
			IotCore: IotCore{
				Enabled: t.Global.IotCoreEnabled.ValueBool(),
			},
			IA: IA{
				Enabled: t.Global.IaEnabled.ValueBool(),
			},
			Gcp: Gcp{
				Project:               t.Global.GcpProject.ValueString(),
				Region:                t.Global.GcpRegion.ValueString(),
				GsmReadServiceAccount: t.Global.GcpGSMServiceAccount.ValueString(),
			},
			Advanced: Advanced{
				PredefinedNamespace: false,
				MemoryStore: MemoryStore{
					Enabled: t.Global.GcpMemoryStoreEnabled.ValueBool(),
				},
				CloudSQL: CloudSQL{
					Enabled: t.Global.GcpCloudSQLEnabled.ValueBool(),
				},
				Secrets: Secrets{
					Manager: "gsm",
				},
			},
			MtlsHAProxy:    false,
			MtlsClearBlade: false,
			GMP:            false,
		},
		Clearblade: Clearblade{
			Replicas:      int(t.Clearblade.BlueReplicas.ValueInt32()),
			GreenReplicas: int(t.Clearblade.GreenReplicas.ValueInt32()),
			License: License{
				Key:        "",
				InstanceID: t.Global.EnterpriseInstanceID.ValueString(),
				AutoRenew: AutoRenew{
					Enabled: true,
				},
			},
			ResourceRequests: ResourceRequests{
				CPU:    int(t.Clearblade.RequestCPU.ValueInt32()),
				Memory: t.Clearblade.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.Clearblade.LimitCPU.ValueInt32()),
				Memory: t.Clearblade.LimitMemory.ValueString(),
			},
			Mqtt: Mqtt{
				AllowDuplicateClientID: t.Clearblade.MqttAllowDuplicateClientID.ValueBool(),
			},
		},
		CbHaproxy: CbHaproxy{
			Enabled:           t.CbHaproxy.Enabled.ValueBool(),
			MonitoringEnabled: false,
			Replicas:          int(t.CbHaproxy.Replicas.ValueInt32()),
			Image:             "haproxy",
			ImageTag:          "2.6-alpine",
			MqttOver443:       t.CbHaproxy.MqttOver443.ValueBool(),
			IP: IP{
				Primary: t.CbHaproxy.PrimaryIP.ValueString(),
				Mqtt:    t.CbHaproxy.MqttIP.ValueString(),
			},
			ResourceRequests: ResourceRequests{
				CPU:    int(t.CbHaproxy.RequestCPU.ValueInt32()),
				Memory: t.CbHaproxy.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbHaproxy.LimitCPU.ValueInt32()),
				Memory: t.CbHaproxy.LimitMemory.ValueString(),
			},
		},
		CbPostgres: CbPostgres{
			Enabled:  t.CbPostgres.Enabled.ValueBool(),
			Image:    "timescale/timescaledb",
			ImageTag: "latest-pg15",
			Replicas: int(t.CbPostgres.Replicas.ValueInt32()),
			ResourceRequests: ResourceRequests{
				CPU:    int(t.CbPostgres.RequestCPU.ValueInt32()),
				Memory: t.CbPostgres.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbPostgres.LimitCPU.ValueInt32()),
				Memory: t.CbPostgres.LimitMemory.ValueString(),
			},
		},
		CbRedis: CbRedis{
			Enabled:  t.CbRedis.Enabled.ValueBool(),
			Image:    "redis",
			ImageTag: "alpine",
			ResourceRequests: ResourceRequests{
				CPU:    int(t.CbRedis.RequestCPU.ValueInt32()),
				Memory: t.CbRedis.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbRedis.LimitCPU.ValueInt32()),
				Memory: t.CbRedis.LimitMemory.ValueString(),
			},
		},
		CbConsole: CbConsole{
			ResourceRequests: ResourceRequests{
				CPU:    int(t.CbConsole.RequestCPU.ValueInt32()),
				Memory: t.CbConsole.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbConsole.LimitCPU.ValueInt32()),
				Memory: t.CbConsole.LimitMemory.ValueString(),
			},
		},
		CbFileHosting: CbFileHosting{
			ResourceRequests: ResourceRequests{
				CPU:    int(t.CbFileHosting.RequestCPU.ValueInt32()),
				Memory: t.CbFileHosting.RequestMemory.ValueString(),
			},
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbFileHosting.LimitCPU.ValueInt32()),
				Memory: t.CbFileHosting.LimitMemory.ValueString(),
			},
		},
		CbIotcore: CbIotcore{
			CheckClearbladeReadiness: false,
			ResourceLimits: ResourceLimits{
				CPU:    int(t.CbIotcore.LimitCPU.ValueInt32()),
				Memory: t.CbIotcore.LimitMemory.ValueString(),
			},
		},
	}
	return h
}
