package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
	EnterpriseGreenVersion    interface{} `yaml:"enterpriseGreenVersion,omitempty"`
	EnterpriseConsoleVersion  interface{} `yaml:"enterpriseConsoleVersion,omitempty"`
	EnterpriseSlot            interface{} `yaml:"enterpriseSlot,omitempty"`
	EnterpriseInstanceID      interface{} `yaml:"enterpriseInstanceID"`
	EnterpriseRegistrationKey interface{} `yaml:"enterpriseRegistrationKey"`
	IotCoreEnabled            bool        `yaml:"iotCoreEnabled"`
	IaEnabled                 bool        `yaml:"iaEnabled"`
	OpsConsoleEnabled         bool        `yaml:"opsConsoleEnabled"`
	GcpCloudSQLEnabled        bool        `yaml:"gcpCloudSQLEnabled"`
	GcpMemoryStoreEnabled     bool        `yaml:"gcpMemoryStoreEnabled"`
	MtlsClearblade            bool        `yaml:"mtlsClearBlade"`
	MtlsHAProxy               bool        `yaml:"mtlsHAProxy"`
	GcpProject                interface{} `yaml:"gcpProject"`
	GcpRegion                 interface{} `yaml:"gcpRegion"`
	GcpGSMServiceAccount      interface{} `yaml:"gcpGSMServiceAccount"`
	SecretManager             string      `yaml:"secretManager"`
	StorageClassName          string      `yaml:"storageClassName"`
}

type Console struct {
	RequestCPU    float32 `yaml:"requestCPU"`
	RequestMemory string  `yaml:"requestMemory"`
	LimitCPU      float32 `yaml:"limitCPU"`
	LimitMemory   string  `yaml:"limitMemory"`
}

type FileHosting struct {
	RequestCPU    float32 `yaml:"requestCPU"`
	RequestMemory string  `yaml:"requestMemory"`
	LimitCPU      float32 `yaml:"limitCPU"`
	LimitMemory   string  `yaml:"limitMemory"`
}

type HAProxy struct {
	Replicas                 int          `yaml:"replicas"`
	RequestCPU               float32      `yaml:"requestCPU"`
	RequestMemory            string       `yaml:"requestMemory"`
	LimitCPU                 float32      `yaml:"limitCPU"`
	LimitMemory              string       `yaml:"limitMemory"`
	Enabled                  bool         `yaml:"enabled"`
	PrimaryIP                interface{}  `yaml:"primaryIP"`
	MqttIP                   interface{}  `yaml:"mqttIP"`
	MqttOver443              bool         `yaml:"mqttOver443"`
	CertRenewal              bool         `yaml:"certRenewal"`
	RenewalDays              int          `yaml:"renewalDays"`
	ControllerVersion        string       `yaml:"controllerVersion"`
	CheckClearbladeReadiness bool         `yaml:"checkClearBladeReadiness"`
	AcmeConfig               []AcmeConfig `yaml:"acmeConfig"`
}

type AcmeConfig struct {
	Directory string   `yaml:"directory"`
	Email     string   `yaml:"email"`
	EabKid    string   `yaml:"eabKid"`
	EabKey    string   `yaml:"eabKey"`
	KeyType   string   `yaml:"keyType"`
	Domains   []string `yaml:"domains"`
	FileName  string   `yaml:"fileName"`
}

type IotCore struct {
	CheckClearbladeReadiness bool    `yaml:"checkClearbladeReadiness"`
	RequestCPU               float32 `yaml:"requestCPU"`
	RequestMemory            string  `yaml:"requestMemory"`
	LimitCPU                 float32 `yaml:"limitCPU"`
	LimitMemory              string  `yaml:"limitMemory"`
	Version                  string  `yaml:"version"`
	Regions                  string  `yaml:"regions"`
}

type Ia struct {
	CheckClearbladeReadiness bool    `yaml:"checkClearbladeReadiness"`
	RequestCPU               float32 `yaml:"requestCPU"`
	RequestMemory            string  `yaml:"requestMemory"`
	LimitCPU                 float32 `yaml:"limitCPU"`
	LimitMemory              string  `yaml:"limitMemory"`
	Version                  string  `yaml:"version"`
}

type Postgres struct {
	Enabled           bool    `yaml:"enabled"`
	Replicas          int     `yaml:"replicas"`
	RequestCPU        float32 `yaml:"requestCPU"`
	RequestMemory     string  `yaml:"requestMemory"`
	LimitCPU          float32 `yaml:"limitCPU"`
	LimitMemory       string  `yaml:"limitMemory"`
	Postgres0DiskName string  `yaml:"postgres0DiskName"`
}

type Redis struct {
	Enabled          bool    `yaml:"enabled"`
	HighAvailability bool    `yaml:"highAvailability"`
	RequestCPU       float32 `yaml:"requestCPU"`
	RequestMemory    string  `yaml:"requestMemory"`
	LimitCPU         float32 `yaml:"limitCPU"`
	LimitMemory      string  `yaml:"limitMemory"`
}

type Clearblade struct {
	BlueReplicas               int               `yaml:"blueReplicas"`
	GreenReplicas              int               `yaml:"greenReplicas"`
	MqttAllowDuplicateClientID bool              `yaml:"mqttAllowDuplicateClientID"`
	RequestCPU                 float32           `yaml:"requestCPU"`
	RequestMemory              string            `yaml:"requestMemory"`
	LimitCPU                   float32           `yaml:"limitCPU"`
	LimitMemory                string            `yaml:"limitMemory"`
	License                    ClearbladeLicense `yaml:"license"`
}

type ClearbladeLicense struct {
	LicenseRenewalWebhooks   []string            `yaml:"renewalWebhooks"`
	MetricsReportingWebhooks []string            `yaml:"metricsWebhooks"`
	AutoRenew                ClearbladeAutoRenew `yaml:"autoRenew"`
}

type ClearbladeAutoRenew struct {
	Enabled bool `yaml:"enabled"`
	Days    int  `yaml:"days"`
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
	EnterpriseGreenVersion    types.String `tfsdk:"enterprise_green_version"`
	EnterpriseConsoleVersion  types.String `tfsdk:"enterprise_console_version"`
	EnterpriseSlot            types.String `tfsdk:"enterprise_slot"`
	EnterpriseInstanceID      types.String `tfsdk:"enterprise_instance_id"`
	EnterpriseRegistrationKey types.String `tfsdk:"enterprise_registration_key"`
	IotCoreEnabled            types.Bool   `tfsdk:"iotcore_enabled"`
	IaEnabled                 types.Bool   `tfsdk:"ia_enabled"`
	OpsConsoleEnabled         types.Bool   `tfsdk:"ops_console_enabled"`
	GcpCloudSQLEnabled        types.Bool   `tfsdk:"gcp_cloudsql_enabled"`
	GcpMemoryStoreEnabled     types.Bool   `tfsdk:"gcp_memorystore_enabled"`
	GcpProject                types.String `tfsdk:"gcp_project"`
	GcpRegion                 types.String `tfsdk:"gcp_region"`
	GcpGSMServiceAccount      types.String `tfsdk:"gcp_gsm_service_account"`
	StorageClassName          types.String `tfsdk:"storage_class_name"`
	MtlsClearblade            types.Bool   `tfsdk:"enable_mtls_clearblade"`
	MtlsHAProxy               types.Bool   `tfsdk:"enable_mtls_haproxy"`
}

type TfConsole struct {
	RequestCPU    types.Float32 `tfsdk:"request_cpu"`
	RequestMemory types.String  `tfsdk:"request_memory"`
	LimitCPU      types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory   types.String  `tfsdk:"limit_memory"`
}

type TfFileHosting struct {
	RequestCPU    types.Float32 `tfsdk:"request_cpu"`
	RequestMemory types.String  `tfsdk:"request_memory"`
	LimitCPU      types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory   types.String  `tfsdk:"limit_memory"`
}

type TfHAProxy struct {
	Replicas                 types.Int32    `tfsdk:"replicas"`
	RequestCPU               types.Float32  `tfsdk:"request_cpu"`
	RequestMemory            types.String   `tfsdk:"request_memory"`
	LimitCPU                 types.Float32  `tfsdk:"limit_cpu"`
	LimitMemory              types.String   `tfsdk:"limit_memory"`
	Enabled                  types.Bool     `tfsdk:"enabled"`
	PrimaryIP                types.String   `tfsdk:"primary_ip"`
	MqttIP                   types.String   `tfsdk:"mqtt_ip"`
	MqttOver443              types.Bool     `tfsdk:"mqtt_over_443"`
	CertRenewal              types.Bool     `tfsdk:"cert_renewal"`
	CheckClearbladeReadiness types.Bool     `tfsdk:"check_clearblade_readiness"`
	RenewalDays              types.Int32    `tfsdk:"renewal_days"`
	ControllerVersion        types.String   `tfsdk:"controller_version"`
	AcmeConfig               []TfAcmeConfig `tfsdk:"acme_config"`
}

type TfAcmeConfig struct {
	Directory types.String `tfsdk:"directory"`
	Email     types.String `tfsdk:"email"`
	EabKid    types.String `tfsdk:"eab_kid"`
	EabKey    types.String `tfsdk:"eab_key"`
	KeyType   types.String `tfsdk:"key_type"`
	Domains   types.List   `tfsdk:"domains"`
	FileName  types.String `tfsdk:"file_name"`
}

type TfIotCore struct {
	CheckClearbladeReadiness types.Bool    `tfsdk:"check_clearblade_readiness"`
	RequestCPU               types.Float32 `tfsdk:"request_cpu"`
	RequestMemory            types.String  `tfsdk:"request_memory"`
	LimitCPU                 types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory              types.String  `tfsdk:"limit_memory"`
	Version                  types.String  `tfsdk:"version"`
	Regions                  types.String  `tfsdk:"regions"`
}

type TfIa struct {
	CheckClearbladeReadiness types.Bool    `tfsdk:"check_clearblade_readiness"`
	RequestCPU               types.Float32 `tfsdk:"request_cpu"`
	RequestMemory            types.String  `tfsdk:"request_memory"`
	LimitCPU                 types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory              types.String  `tfsdk:"limit_memory"`
	Version                  types.String  `tfsdk:"version"`
}

type TfPostgres struct {
	Enabled           types.Bool    `tfsdk:"enabled"`
	Replicas          types.Int32   `tfsdk:"replicas"`
	RequestCPU        types.Float32 `tfsdk:"request_cpu"`
	RequestMemory     types.String  `tfsdk:"request_memory"`
	LimitCPU          types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory       types.String  `tfsdk:"limit_memory"`
	Postgres0DiskName types.String  `tfsdk:"postgres0_disk_name"`
}

type TfRedis struct {
	Enabled          types.Bool    `tfsdk:"enabled"`
	HighAvailability types.Bool    `tfsdk:"high_availability"`
	RequestCPU       types.Float32 `tfsdk:"request_cpu"`
	RequestMemory    types.String  `tfsdk:"request_memory"`
	LimitCPU         types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory      types.String  `tfsdk:"limit_memory"`
}

type TfClearblade struct {
	BlueReplicas               types.Int32   `tfsdk:"blue_replicas"`
	GreenReplicas              types.Int32   `tfsdk:"green_replicas"`
	MqttAllowDuplicateClientID types.Bool    `tfsdk:"mqtt_allow_duplicate_client_id"`
	LicenseRenewalWebhooks     types.List    `tfsdk:"license_renewal_webhooks"`
	MetricsReportingWebhooks   types.List    `tfsdk:"metrics_reporting_webhooks"`
	RequestCPU                 types.Float32 `tfsdk:"request_cpu"`
	RequestMemory              types.String  `tfsdk:"request_memory"`
	LimitCPU                   types.Float32 `tfsdk:"limit_cpu"`
	LimitMemory                types.String  `tfsdk:"limit_memory"`
}

func (t *TfHelmValues) toHelmValues() (*HelmValues, diag.Diagnostics) {
	var licenseRenewalWebhooks []string
	if diags := t.Clearblade.LicenseRenewalWebhooks.ElementsAs(context.Background(), &licenseRenewalWebhooks, false); diags.HasError() {
		return nil, diags
	}

	var metricsReportingWebhooks []string
	if diags := t.Clearblade.MetricsReportingWebhooks.ElementsAs(context.Background(), &metricsReportingWebhooks, false); diags.HasError() {
		return nil, diags
	}

	acmeConfigs, diags := parseAcmeConfigs(t)
	if diags != nil && diags.HasError() {
		return nil, diags
	}

	h := &HelmValues{
		Global: Global{
			Cloud:                     "gcp",
			Namespace:                 t.Global.Namespace.ValueString(),
			ImagePullerSecret:         t.Global.ImagePullerSecret.ValueString(),
			EnterpriseBaseURL:         t.Global.EnterpriseBaseURL.ValueString(),
			EnterpriseBlueVersion:     t.Global.EnterpriseBlueVersion.ValueString(),
			EnterpriseGreenVersion:    t.Global.EnterpriseGreenVersion.ValueString(),
			EnterpriseConsoleVersion:  t.Global.EnterpriseConsoleVersion.ValueString(),
			EnterpriseSlot:            t.Global.EnterpriseSlot.ValueString(),
			EnterpriseInstanceID:      t.Global.EnterpriseInstanceID.ValueString(),
			EnterpriseRegistrationKey: t.Global.EnterpriseRegistrationKey.ValueString(),
			IotCoreEnabled:            t.Global.IotCoreEnabled.ValueBool(),
			IaEnabled:                 t.Global.IaEnabled.ValueBool(),
			OpsConsoleEnabled:         t.Global.OpsConsoleEnabled.ValueBool(),
			GcpCloudSQLEnabled:        t.Global.GcpCloudSQLEnabled.ValueBool(),
			GcpMemoryStoreEnabled:     t.Global.GcpMemoryStoreEnabled.ValueBool(),
			MtlsClearblade:            t.Global.MtlsClearblade.ValueBool(),
			MtlsHAProxy:               t.Global.MtlsHAProxy.ValueBool(),
			GcpProject:                t.Global.GcpProject.ValueString(),
			GcpRegion:                 t.Global.GcpRegion.ValueString(),
			GcpGSMServiceAccount:      t.Global.GcpGSMServiceAccount.ValueString(),
			SecretManager:             "gsm",
			StorageClassName:          t.Global.StorageClassName.ValueString(),
		},
		CbConsole: Console{
			RequestCPU:    t.CbConsole.RequestCPU.ValueFloat32(),
			RequestMemory: t.CbConsole.RequestMemory.ValueString(),
			LimitCPU:      t.CbConsole.LimitCPU.ValueFloat32(),
			LimitMemory:   t.CbConsole.LimitMemory.ValueString(),
		},
		CbFileHosting: FileHosting{
			RequestCPU:    t.CbFileHosting.RequestCPU.ValueFloat32(),
			RequestMemory: t.CbFileHosting.RequestMemory.ValueString(),
			LimitCPU:      t.CbFileHosting.LimitCPU.ValueFloat32(),
			LimitMemory:   t.CbFileHosting.LimitMemory.ValueString(),
		},
		CbHaproxy: HAProxy{
			Replicas:                 int(t.CbHaproxy.Replicas.ValueInt32()),
			RequestCPU:               t.CbHaproxy.RequestCPU.ValueFloat32(),
			RequestMemory:            t.CbHaproxy.RequestMemory.ValueString(),
			LimitCPU:                 t.CbHaproxy.LimitCPU.ValueFloat32(),
			LimitMemory:              t.CbHaproxy.LimitMemory.ValueString(),
			Enabled:                  t.CbHaproxy.Enabled.ValueBool(),
			PrimaryIP:                t.CbHaproxy.PrimaryIP.ValueString(),
			MqttIP:                   t.CbHaproxy.MqttIP.ValueString(),
			MqttOver443:              t.CbHaproxy.MqttOver443.ValueBool(),
			CertRenewal:              t.CbHaproxy.CertRenewal.ValueBool(),
			CheckClearbladeReadiness: t.CbHaproxy.CheckClearbladeReadiness.ValueBool(),
			RenewalDays:              int(t.CbHaproxy.RenewalDays.ValueInt32()),
			ControllerVersion:        t.CbHaproxy.ControllerVersion.ValueString(),
			AcmeConfig:               acmeConfigs,
		},
		CbIotcore: IotCore{
			CheckClearbladeReadiness: t.CbIotcore.CheckClearbladeReadiness.ValueBool(),
			RequestCPU:               t.CbIotcore.RequestCPU.ValueFloat32(),
			RequestMemory:            t.CbIotcore.RequestMemory.ValueString(),
			LimitCPU:                 t.CbIotcore.LimitCPU.ValueFloat32(),
			LimitMemory:              t.CbIotcore.LimitMemory.ValueString(),
		},
		CbIa: Ia{
			CheckClearbladeReadiness: t.CbIa.CheckClearbladeReadiness.ValueBool(),
			RequestCPU:               t.CbIa.RequestCPU.ValueFloat32(),
			RequestMemory:            t.CbIa.RequestMemory.ValueString(),
			LimitCPU:                 t.CbIa.LimitCPU.ValueFloat32(),
			LimitMemory:              t.CbIa.LimitMemory.ValueString(),
			Version:                  t.CbIa.Version.ValueString(),
		},
		CbPostgres: Postgres{
			Enabled:           t.CbPostgres.Enabled.ValueBool(),
			Replicas:          int(t.CbPostgres.Replicas.ValueInt32()),
			RequestCPU:        t.CbPostgres.RequestCPU.ValueFloat32(),
			RequestMemory:     t.CbPostgres.RequestMemory.ValueString(),
			LimitCPU:          t.CbPostgres.LimitCPU.ValueFloat32(),
			LimitMemory:       t.CbPostgres.LimitMemory.ValueString(),
			Postgres0DiskName: t.CbPostgres.Postgres0DiskName.ValueString(),
		},
		CbRedis: Redis{
			Enabled:          t.CbRedis.Enabled.ValueBool(),
			HighAvailability: t.CbRedis.HighAvailability.ValueBool(),
			RequestCPU:       t.CbRedis.RequestCPU.ValueFloat32(),
			RequestMemory:    t.CbRedis.RequestMemory.ValueString(),
			LimitCPU:         t.CbRedis.LimitCPU.ValueFloat32(),
			LimitMemory:      t.CbRedis.LimitMemory.ValueString(),
		},
		Clearblade: Clearblade{
			License: ClearbladeLicense{
				LicenseRenewalWebhooks:   licenseRenewalWebhooks,
				MetricsReportingWebhooks: metricsReportingWebhooks,
				AutoRenew: ClearbladeAutoRenew{
					Enabled: true,
					Days:    15,
				},
			},
			BlueReplicas:               int(t.Clearblade.BlueReplicas.ValueInt32()),
			GreenReplicas:              int(t.Clearblade.GreenReplicas.ValueInt32()),
			MqttAllowDuplicateClientID: t.Clearblade.MqttAllowDuplicateClientID.ValueBool(),
			RequestCPU:                 t.Clearblade.RequestCPU.ValueFloat32(),
			RequestMemory:              t.Clearblade.RequestMemory.ValueString(),
			LimitCPU:                   t.Clearblade.LimitCPU.ValueFloat32(),
			LimitMemory:                t.Clearblade.LimitMemory.ValueString(),
		},
	}
	return h, nil
}

func parseAcmeConfigs(tf *TfHelmValues) ([]AcmeConfig, diag.Diagnostics) {
	configs := tf.CbHaproxy.AcmeConfig
	acmeConfigs := make([]AcmeConfig, len(configs))
	for i, acmeConfig := range configs {
		var domains []string
		if diags := acmeConfig.Domains.ElementsAs(context.Background(), &domains, false); diags.HasError() {
			return nil, diags
		}

		acmeConfigs[i] = AcmeConfig{
			Directory: acmeConfig.Directory.ValueString(),
			Email:     acmeConfig.Email.ValueString(),
			EabKid:    acmeConfig.EabKid.ValueString(),
			EabKey:    acmeConfig.EabKey.ValueString(),
			KeyType:   acmeConfig.KeyType.ValueString(),
			Domains:   domains,
			FileName:  acmeConfig.FileName.ValueString(),
		}
	}

	return acmeConfigs, nil
}
