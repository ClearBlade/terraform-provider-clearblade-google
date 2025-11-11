package provider

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TLSCertificateResource{}
var _ resource.ResourceWithImportState = &TLSCertificateResource{}

const expiredTLSCert = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCTtXeP+g+dtmuF
F86q2PmqvZz84hm1EcDRRGmzqBeC/VPKVRD9NvMvEikKtJN8w9zI1WaMkzgBfA/J
umd+enA4o9p4MGmVY+u8cn7lmTuN/MhyzCfU/a3Uc7Kw6UysRJdPhSR866vW3dDP
pJdVfZ0LwCgu7C0eWrzKMzJCDckDCVz4zlUKtmMUipU17d+3XEE4P97TBQtOh2YC
YeAAkNJZB0SmpHePEUS1np+2FG8swXLu8TQ0Bc9Kdz475LKjvmKLum6jufPZdxID
4yYL5M/0y4Wy3+Fj4l3+VCWcasiEAz2C3DWHkDthMKh+kFkXgpEr55tTS4xpx7gt
SOGFLB3FAgMBAAECggEAAfonnc6B14FZ2pf5PPm3C9VKbOsk33LboyF4jb5WBDua
W8a8Obt60Vo7oOhOYhjoE2sh2odc9E4iEvfzCzMd3fA5jCrPuv9xqB3bO30L2kh0
MW8wqE31/fZHgc05qMOpR9f3J1HrRK7G7QSdvvf1unLJxukD6Jhb1xQM2+v90RPR
ivIkl8tjpr7Q2yfVXjOk1WqzOesbeGjFIJZwsaD87v773qnkP/HWfa7vhuKsMCTh
F9V56pWnga9/jzzNVZ1aysof3vIPuH6BVdchwvdyBXV3rmTaLRWIORF9so4pGhdc
xhq0Bd/TONBB/qwmWXV1yLth4xRxu8UvxhKwl+GyuQKBgQDOZBs0+lyxOTOb+PPN
JxBFTxvqAE884DxJjP2DM0QSO1QpMG2hfsiE4y/KfiHCcjE5AnuwMJxBBza5ZNmA
VUBdenImWNMmrWyQh0V/NC5BP8ZHf9FFbNU63spe9mlUOCW6M1SW+Rha6jDRycPn
do+DV+Oktma/UvS19ZFYQ7+nvQKBgQC3NnXU3YF4nt6RW7VqLDLgiOcR5czR2tqH
0TkXQWQ5TKJLNFUQnxrhFVq1dlbkz9i/Xp0hXC2iAeH4Z8EgPO2FypzvxwSqIyXE
k+XrrwRCIczOw6UFxMweD1pHCUQplvJYNC2g3lNfAHP3aRoP1kMpNAyJjkdrDDIz
0wygXI0KqQKBgDSm5yXxxNnVXOwqa4/nqkf0MYvVvmEqV1bwJ/BjkLcR+Zt0ZlNv
s5nrF1MSMGyZkyMXFhTRodsZCwXqy23o0b3HMf3EZUGVtn98cudLmY09xsiQvAN5
C0C0e24UcLRyinVhCPBm5aaz3fZ3AYo3/c6lCkcH3Vhrwk/1MLoStn8FAoGAHqRk
Jrr5WYQws10ERYKo67bZ9rtZe0vAOvD6PHJ6Yb74cd2J0KLbqwOYTTtCozhEBxW6
8AZrt2nbMmGgAlVOYI8Xml7N3+rK+UrHLJjz/F+M0pQUOJfGj8x/i9v344DUfX4U
l5A3n28C9kFE05tBVlXXNvZt6XB7wQEuXm+8QykCgYEAmpWeDnU6lc3zMpYxigab
t+msCNC/ludy5xG6XBb6DhPf11DvLmK1hI7XNVjkskNt1+75IpjEmcfs+sWTr67V
RVpYP2JJCxcxxedaxWikBKx89BuP4+MqoOVsrt80uZyKyEqSWTuPTL7H/8Au3K23
uLXanikWAHV/s8yHW04p1Fo=
-----END PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIDOTCCAiGgAwIBAgIUYAy38A3jZIfITZoywSj0TPDbpFswDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCVVMxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNTAxMTQwNTA3MjBaFw0yNTAx
MTUwNTA3MjBaMEUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQCTtXeP+g+dtmuFF86q2PmqvZz84hm1EcDRRGmzqBeC
/VPKVRD9NvMvEikKtJN8w9zI1WaMkzgBfA/Jumd+enA4o9p4MGmVY+u8cn7lmTuN
/MhyzCfU/a3Uc7Kw6UysRJdPhSR866vW3dDPpJdVfZ0LwCgu7C0eWrzKMzJCDckD
CVz4zlUKtmMUipU17d+3XEE4P97TBQtOh2YCYeAAkNJZB0SmpHePEUS1np+2FG8s
wXLu8TQ0Bc9Kdz475LKjvmKLum6jufPZdxID4yYL5M/0y4Wy3+Fj4l3+VCWcasiE
Az2C3DWHkDthMKh+kFkXgpEr55tTS4xpx7gtSOGFLB3FAgMBAAGjITAfMB0GA1Ud
DgQWBBQC6cKrPuSsK7Non0yS9V5+TLFluzANBgkqhkiG9w0BAQsFAAOCAQEAEyxb
8l2x66UH+P8pmW4JcVZ1TOKl6F1WRNanYMKTM+YJQ3F4r4gI6ka29rF2JDemBbtk
KmHiNz/ZW1K2sDPyDjpapKEn5rTg9eVJUjWIfcEzWHXE1kjQ1i5HQOXKTi0wylIg
AxXx/4cWwOwJfCvPPeZkLQ1MpM/GFmhlg0lKf88BBGWQN+H3Kc4mxX6yfcdPQcmn
Lj7rxN6KGPReEyHvQCYLXYHkQCvNHUXeCw/7V8fAqBbERIswLb18nCH7Dmp7Wv9C
KP8VRJFHbBKSXxAHsi4uRDmwup1oVjnkU+i+F9H6Cj3MtCrFGNyQbxiVmeTET35R
KXTgtaWNL/Pv4fWEdg==
-----END CERTIFICATE-----`

func NewTLSCertificateResource() resource.Resource {
	return &TLSCertificateResource{}
}

// TLSCertificateResource defines the resource implementation.
type TLSCertificateResource struct {
	client *secretmanager.Client
}

// TLSCertificateResourceModel describes the resource data model.
type TLSCertificateResourceModel struct {
	ProjectId       types.String `tfsdk:"project_id"`
	Namespace       types.String `tfsdk:"namespace"`
	Suffix          types.String `tfsdk:"suffix"`
	TLSCertificates types.Map    `tfsdk:"tls_certificates"`
	SecretId        types.String `tfsdk:"secret_id"`
}

func (t *TLSCertificateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tls_certificate"
}

func (t *TLSCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Create and expired TLS certificate and store it in GCP Secrets",

		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				MarkdownDescription: "GCP project Id for storing MEK",
				Required:            true,
			},
			"namespace": schema.StringAttribute{
				MarkdownDescription: "Instance namespace",
				Required:            true,
			},
			"suffix": schema.StringAttribute{
				MarkdownDescription: "Secret Id suffix",
				Required:            true,
			},
			"tls_certificates": schema.MapAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Map of certificate names to PEM encoded certificate strings",
				Required:            true,
			},
			"secret_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (t *TLSCertificateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*secretmanager.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *secretmanager.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	t.client = client
}

func (t *TLSCertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TLSCertificateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	if err := createSecret(ctx, t.client, data.ProjectId.ValueString(), secretId); err != nil {
		resp.Diagnostics.AddError("Failed to create secret", err.Error())
		return
	}
	data.SecretId = types.StringValue(secretId)
	certBytes, err := data.getSecretBytes()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get secret bytes", err.Error())
		return
	}

	if err := addSecretVersion(ctx, t.client, data.ProjectId.ValueString(), secretId, certBytes); err != nil {
		resp.Diagnostics.AddError("Failed to create tls certificate secret", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (t *TLSCertificateResourceModel) getSecretBytes() ([]byte, error) {
	certs := map[string]string{}
	for key, value := range t.TLSCertificates.Elements() {
		strValue, ok := value.(types.String)
		if !ok {
			return nil, fmt.Errorf("value is not a string, is: %T", value)
		}

		contents := strValue.ValueString()
		base64Contents := base64.StdEncoding.EncodeToString([]byte(contents))
		certs[key] = base64Contents
	}

	if len(certs) == 0 {
		// Put in a default cert so that HaProxy can start
		certs["clearblade-0.pem"] = expiredTLSCert
	}

	certBytes, err := json.Marshal(certs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal certificates: %w", err)
	}

	return certBytes, nil
}

func (t *TLSCertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TLSCertificateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	resource := getSecretResourceName(data.ProjectId.ValueString(), secretId) + "/versions/latest"
	secReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: resource,
	}
	rval, err := t.client.AccessSecretVersion(ctx, secReq)
	if err != nil {
		resp.Diagnostics.AddError("Faled to get TLS certificate secret", err.Error())
		return
	}
	if rval.Payload.Data == nil || string(rval.Payload.Data) == "" {
		resp.Diagnostics.AddError("Faled to get TLS certificate secret data", "Empty payload")
	}
	data.SecretId = types.StringValue(secretId)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (t *TLSCertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TLSCertificateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	certBytes, err := data.getSecretBytes()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get secret bytes", err.Error())
		return
	}

	if err := addSecretVersion(ctx, t.client, data.ProjectId.ValueString(), secretId, certBytes); err != nil {
		resp.Diagnostics.AddError("Failed to update tls certificate", err.Error())
		return
	}
	data.SecretId = types.StringValue(secretId)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (t *TLSCertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TLSCertificateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	resource := getSecretResourceName(data.ProjectId.ValueString(), secretId)
	delReq := &secretmanagerpb.DeleteSecretRequest{
		Name: resource,
	}
	if err := t.client.DeleteSecret(ctx, delReq); err != nil {
		resp.Diagnostics.AddError("Failed to delete TLS certificate", err.Error())
		return
	}
}

func (e *TLSCertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
