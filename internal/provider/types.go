package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type TfHelmValues struct {
	Global        TfGlobal      `tfsdk:"global"`
	CbConsole     TfConsole     `tfsdk:"cb-console"`
	CbFileHosting TfFileHosting `tfsdk:"cb-file-hosting"`
	CbHaproxy     TfHAProxy     `tfsdk:"cb-haproxy"`
	CbIotcore     TfIotCore     `tfsdk:"cb-iotcore"`
	CbIa          TfIa          `tfsdk:"cb-ia"`
	CbPostgres    TfPostgres    `tfsdk:"cb-postgres"`
	CbRedis       TfRedis       `tfsdk:"cb-redis"`
	Clearblade    TfClearblade  `tfsdk:"clearblade"`
}

type TfGlobal struct {
	Cloud                     types.String `tfsdk:"cloud"`
	Namespace                 types.String `tfsdk:"namespace"`
	ImagePullerSecret         types.String `tfsdk:"imagePullerSecret"`
	EnterpriseBaseURL         types.String `tfsdk:"enterpriseBaseURL"`
	EnterpriseBlueVersion     types.String `tfsdk:"enterpriseBlueVersion"`
	EnterpriseInstanceID      types.String `tfsdk:"enterpriseInstanceID"`
	EnterpriseRegistrationKey types.String `tfsdk:"enterpriseRegistrationKey"`
	IotCoreEnabled            types.Bool   `tfsdk:"iotCoreEnabled"`
	IaEnabled                 types.Bool   `tfsdk:"iaEnabled"`
	GcpCloudSQLEnabled        types.Bool   `tfsdk:"gcpCloudSQLEnabled"`
	GcpMemoryStoreEnabled     types.Bool   `tfsdk:"gcpMemoryStoreEnabled"`
	GcpProject                types.String `tfsdk:"gcpProject"`
	GcpRegion                 types.String `tfsdk:"gcpRegion"`
	GcpGSMServiceAccount      types.String `tfsdk:"gcpGSMServiceAccount"`
	StorageClassName          types.String `tfsdk:"storageClassName"`
}

type TfConsole struct {
	RequestCPU    types.Int32  `tfsdk:"requestCPU"`
	RequestMemory types.String `tfsdk:"requestMemory"`
	LimitCPU      types.Int32  `tfsdk:"limitCPU"`
	LimitMemory   types.String `tfsdk:"limitMemory"`
}

type TfFileHosting struct {
	RequestCPU    types.Int32  `tfsdk:"requestCPU"`
	RequestMemory types.String `tfsdk:"requestMemory"`
	LimitCPU      types.Int32  `tfsdk:"limitCPU"`
	LimitMemory   types.String `tfsdk:"limitMemory"`
}

type TfHAProxy struct {
	Replicas      types.Int32  `tfsdk:"replicas"`
	RequestCPU    types.Int32  `tfsdk:"requestCPU"`
	RequestMemory types.String `tfsdk:"requestMemory"`
	LimitCPU      types.Int32  `tfsdk:"limitCPU"`
	LimitMemory   types.String `tfsdk:"limitMemory"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	PrimaryIP     types.String `tfsdk:"primaryIP"`
	MqttIP        types.String `tfsdk:"mqttIP"`
	MqttOver443   types.Bool   `tfsdk:"mqttOver443"`
}

type TfIotCore struct {
	CheckClearbladeReadiness types.Bool   `tfsdk:"checkClearbladeReadiness"`
	RequestCPU               types.Int32  `tfsdk:"requestCPU"`
	RequestMemory            types.String `tfsdk:"requestMemory"`
	LimitCPU                 types.Int32  `tfsdk:"limitCPU"`
	LimitMemory              types.String `tfsdk:"limitMemory"`
}

type TfIa struct {
	CheckClearbladeReadiness types.Bool   `tfsdk:"checkClearbladeReadiness"`
	RequestCPU               types.Int32  `tfsdk:"requestCPU"`
	RequestMemory            types.String `tfsdk:"requestMemory"`
	LimitCPU                 types.Int32  `tfsdk:"limitCPU"`
	LimitMemory              types.String `tfsdk:"limitMemory"`
}

type TfPostgres struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	Replicas          types.Int32  `tfsdk:"replicas"`
	RequestCPU        types.Int32  `tfsdk:"requestCPU"`
	RequestMemory     types.String `tfsdk:"requestMemory"`
	LimitCPU          types.Int32  `tfsdk:"limitCPU"`
	LimitMemory       types.String `tfsdk:"limitMemory"`
	Postgres0DiskName types.String `tfsdk:"postgres0DiskName"`
}

type TfRedis struct {
	Enabled          types.Bool   `tfsdk:"enabled"`
	HighAvailability types.Bool   `tfsdk:"highAvailability"`
	RequestCPU       types.Int32  `tfsdk:"requestCPU"`
	RequestMemory    types.String `tfsdk:"requestMemory"`
	LimitCPU         types.Int32  `tfsdk:"limitCPU"`
	LimitMemory      types.String `tfsdk:"limitMemory"`
}

type TfClearblade struct {
	BlueReplicas               types.Int32  `tfsdk:"blueReplicas"`
	GreenReplicas              types.Int32  `tfsdk:"greenReplicas"`
	MqttAllowDuplicateClientID types.Bool   `tfsdk:"mqttAllowDuplicateClientID"`
	RequestCPU                 types.Int32  `tfsdk:"requestCPU"`
	RequestMemory              types.String `tfsdk:"requestMemory"`
	LimitCPU                   types.Int32  `tfsdk:"limitCPU"`
	LimitMemory                types.String `tfsdk:"limitMemory"`
}
