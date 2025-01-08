package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type HelmValues struct {
	Global        Global      `yaml:"global"`
	CbConsole     Console     `yaml:"cb-console"`
	CbFileHosting FileHosting `yaml:"cb-file-hosting"`
	CbHaproxy     HAProxy     `yaml:"cb-haproxy"`
	CbIotcore     IotCore     `yaml:"cb-iotcore"`
	CbIa          Ia          `yaml:"cb-ia"`
	CbPostgres    Postgres    `yaml:"cb-postgres"`
	CbRedis       Redis       `yaml:"cb-redis"`
	Clearblade    Clearblade  `yaml:"clearblade"`
}

type Global struct {
	Cloud                     string      `yaml:"cloud"`
	Namespace                 interface{} `yaml:"namespace"`
	ImagePullerSecret         interface{} `yaml:"imagePullerSecret"`
	EnterpriseBaseURL         interface{} `yaml:"enterpriseBaseURL"`
	EnterpriseBlueVersion     interface{} `yaml:"enterpriseBlueVersion"`
	EnterpriseInstanceID      interface{} `yaml:"enterpriseInstanceID"`
	EnterpriseRegistrationKey interface{} `yaml:"enterpriseRegistrationKey"`
	IotCoreEnabled            bool        `yaml:"iotCoreEnabled"`
	IaEnabled                 bool        `yaml:"iaEnabled"`
	GcpCloudSQLEnabled        bool        `yaml:"gcpCloudSQLEnabled"`
	GcpMemoryStoreEnabled     bool        `yaml:"gcpMemoryStoreEnabled"`
	GcpProject                interface{} `yaml:"gcpProject"`
	GcpRegion                 interface{} `yaml:"gcpRegion"`
	GcpGSMServiceAccount      interface{} `yaml:"gcpGSMServiceAccount"`
	SecretManager             string      `yaml:"secretManager"`
	StorageClassName          string      `yaml:"storageClassName"`
}

type Console struct {
	RequestCPU    int    `yaml:"requestCPU"`
	RequestMemory string `yaml:"requestMemory"`
	LimitCPU      int    `yaml:"limitCPU"`
	LimitMemory   string `yaml:"limitMemory"`
}

type FileHosting struct {
	RequestCPU    int    `yaml:"requestCPU"`
	RequestMemory string `yaml:"requestMemory"`
	LimitCPU      int    `yaml:"limitCPU"`
	LimitMemory   string `yaml:"limitMemory"`
}

type HAProxy struct {
	Replicas      int         `yaml:"replicas"`
	RequestCPU    int         `yaml:"requestCPU"`
	RequestMemory string      `yaml:"requestMemory"`
	LimitCPU      int         `yaml:"limitCPU"`
	LimitMemory   string      `yaml:"limitMemory"`
	Enabled       bool        `yaml:"enabled"`
	PrimaryIP     interface{} `yaml:"primaryIP"`
	MqttIP        interface{} `yaml:"mqttIP"`
	MqttOver443   bool        `yaml:"mqttOver443"`
}

type IotCore struct {
	CheckClearbladeReadiness bool   `yaml:"checkClearbladeReadiness"`
	RequestCPU               int    `yaml:"requestCPU"`
	RequestMemory            string `yaml:"requestMemory"`
	LimitCPU                 int    `yaml:"limitCPU"`
	LimitMemory              string `yaml:"limitMemory"`
}

type Ia struct {
	CheckClearbladeReadiness bool   `yaml:"checkClearbladeReadiness"`
	RequestCPU               int    `yaml:"requestCPU"`
	RequestMemory            string `yaml:"requestMemory"`
	LimitCPU                 int    `yaml:"limitCPU"`
	LimitMemory              string `yaml:"limitMemory"`
}

type Postgres struct {
	Enabled           bool   `yaml:"enabled"`
	Replicas          int    `yaml:"replicas"`
	RequestCPU        int    `yaml:"requestCPU"`
	RequestMemory     string `yaml:"requestMemory"`
	LimitCPU          int    `yaml:"limitCPU"`
	LimitMemory       string `yaml:"limitMemory"`
	Postgres0DiskName string `yaml:"postgres0DiskName"`
}

type Redis struct {
	Enabled          bool   `yaml:"enabled"`
	HighAvailability bool   `yaml:"highAvailability"`
	RequestCPU       int    `yaml:"requestCPU"`
	RequestMemory    string `yaml:"requestMemory"`
	LimitCPU         int    `yaml:"limitCPU"`
	LimitMemory      string `yaml:"limitMemory"`
}

type Clearblade struct {
	BlueReplicas               int    `yaml:"blueReplicas"`
	GreenReplicas              int    `yaml:"greenReplicas"`
	MqttAllowDuplicateClientID bool   `yaml:"mqttAllowDuplicateClientID"`
	RequestCPU                 int    `yaml:"requestCPU"`
	RequestMemory              string `yaml:"requestMemory"`
	LimitCPU                   int    `yaml:"limitCPU"`
	LimitMemory                string `yaml:"limitMemory"`
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
			Cloud:                     "gcp",
			Namespace:                 t.Global.Namespace.ValueString(),
			ImagePullerSecret:         t.Global.ImagePullerSecret.ValueString(),
			EnterpriseBaseURL:         t.Global.EnterpriseBaseURL.ValueString(),
			EnterpriseBlueVersion:     t.Global.EnterpriseBlueVersion.ValueString(),
			EnterpriseInstanceID:      t.Global.EnterpriseInstanceID.ValueString(),
			EnterpriseRegistrationKey: t.Global.EnterpriseRegistrationKey.ValueString(),
			IotCoreEnabled:            t.Global.IotCoreEnabled.ValueBool(),
			IaEnabled:                 t.Global.IaEnabled.ValueBool(),
			GcpCloudSQLEnabled:        t.Global.GcpCloudSQLEnabled.ValueBool(),
			GcpMemoryStoreEnabled:     t.Global.GcpMemoryStoreEnabled.ValueBool(),
			GcpProject:                t.Global.GcpProject.ValueString(),
			GcpRegion:                 t.Global.GcpRegion.ValueString(),
			GcpGSMServiceAccount:      t.Global.GcpGSMServiceAccount.ValueString(),
			SecretManager:             "gsm",
			StorageClassName:          t.Global.StorageClassName.ValueString(),
		},
		CbConsole: Console{
			RequestCPU:    int(t.CbConsole.RequestCPU.ValueInt32()),
			RequestMemory: t.CbConsole.RequestMemory.ValueString(),
			LimitCPU:      int(t.CbConsole.LimitCPU.ValueInt32()),
			LimitMemory:   t.CbConsole.LimitMemory.ValueString(),
		},
		CbFileHosting: FileHosting{
			RequestCPU:    int(t.CbFileHosting.RequestCPU.ValueInt32()),
			RequestMemory: t.CbFileHosting.RequestMemory.ValueString(),
			LimitCPU:      int(t.CbFileHosting.LimitCPU.ValueInt32()),
			LimitMemory:   t.CbFileHosting.LimitMemory.ValueString(),
		},
		CbHaproxy: HAProxy{
			Replicas:      int(t.CbHaproxy.Replicas.ValueInt32()),
			RequestCPU:    int(t.CbHaproxy.RequestCPU.ValueInt32()),
			RequestMemory: t.CbHaproxy.RequestMemory.ValueString(),
			LimitCPU:      int(t.CbHaproxy.LimitCPU.ValueInt32()),
			LimitMemory:   t.CbHaproxy.LimitMemory.ValueString(),
			Enabled:       t.CbHaproxy.Enabled.ValueBool(),
			PrimaryIP:     t.CbHaproxy.PrimaryIP.ValueString(),
			MqttIP:        t.CbHaproxy.MqttIP.ValueString(),
			MqttOver443:   t.CbHaproxy.MqttOver443.ValueBool(),
		},
		CbIotcore: IotCore{
			CheckClearbladeReadiness: t.CbIotcore.CheckClearbladeReadiness.ValueBool(),
			RequestCPU:               int(t.CbIotcore.RequestCPU.ValueInt32()),
			RequestMemory:            t.CbIotcore.RequestMemory.ValueString(),
			LimitCPU:                 int(t.CbIotcore.LimitCPU.ValueInt32()),
			LimitMemory:              t.CbIotcore.LimitMemory.ValueString(),
		},
		CbIa: Ia{
			CheckClearbladeReadiness: t.CbIa.CheckClearbladeReadiness.ValueBool(),
			RequestCPU:               int(t.CbIa.RequestCPU.ValueInt32()),
			RequestMemory:            t.CbIa.RequestMemory.ValueString(),
			LimitCPU:                 int(t.CbIa.LimitCPU.ValueInt32()),
			LimitMemory:              t.CbIa.LimitMemory.ValueString(),
		},
		CbPostgres: Postgres{
			Enabled:           t.CbPostgres.Enabled.ValueBool(),
			Replicas:          int(t.CbPostgres.Replicas.ValueInt32()),
			RequestCPU:        int(t.CbPostgres.RequestCPU.ValueInt32()),
			RequestMemory:     t.CbPostgres.RequestMemory.ValueString(),
			LimitCPU:          int(t.CbPostgres.LimitCPU.ValueInt32()),
			LimitMemory:       t.CbPostgres.LimitMemory.ValueString(),
			Postgres0DiskName: t.CbPostgres.Postgres0DiskName.ValueString(),
		},
		CbRedis: Redis{
			Enabled:          t.CbRedis.Enabled.ValueBool(),
			HighAvailability: t.CbRedis.HighAvailability.ValueBool(),
			RequestCPU:       int(t.CbRedis.RequestCPU.ValueInt32()),
			RequestMemory:    t.CbRedis.RequestMemory.ValueString(),
			LimitCPU:         int(t.CbRedis.LimitCPU.ValueInt32()),
			LimitMemory:      t.CbRedis.LimitMemory.ValueString(),
		},
		Clearblade: Clearblade{
			BlueReplicas:               int(t.Clearblade.BlueReplicas.ValueInt32()),
			GreenReplicas:              int(t.Clearblade.GreenReplicas.ValueInt32()),
			MqttAllowDuplicateClientID: t.Clearblade.MqttAllowDuplicateClientID.ValueBool(),
			RequestCPU:                 int(t.Clearblade.RequestCPU.ValueInt32()),
			RequestMemory:              t.Clearblade.RequestMemory.ValueString(),
			LimitCPU:                   int(t.Clearblade.LimitCPU.ValueInt32()),
			LimitMemory:                t.Clearblade.LimitMemory.ValueString(),
		},
	}
	return h
}
