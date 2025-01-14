package provider

import (
	"context"
	"fmt"
	"strings"

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

const expiredTLSCert = `-----BEGIN CERTIFICATE-----
MIIFazCCA1OgAwIBAgIUGFJh3lYf5VnghjyBYpZsbkuJXiEwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCVVMxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yNTAxMTMxODE0MjZaFw0yNTAx
MTQxODE0MjZaMEUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggIiMA0GCSqGSIb3DQEB
AQUAA4ICDwAwggIKAoICAQDCvzepJx89Av8hGS1ZfKE6+psAB9iMz2iGLy5+kYW0
ih/iN9N+8/SzqvCvtxswd4H8EYgxfz96vE9rH+EicueyfVNlQLf4oF0TN0751/+B
P+mSudd9A6XYSRSw3eOlOLedGuYpcMvz/cW8tpjVWJLvCPzj6uws8UMpW2NjI6rr
ZwwDwohJXmrmaj6KGStyCxbs2pHSyLpnLvJ4cxAPkkuXzECUG0E8T01T8zAyCPE3
Fw6EPlnQoiqPk6klsqxSnjAD48DhSxDTHsn+Ss2I9CqLrPQwqWFdzaaOqSnosKgO
0i+m4VLjAkJ9NDjvPhOAW76m5sECjYg2chy86WEImQOxjfqBQ0+NfHDnsCpi9BkS
wK0eAUUon9WcRwiHGw1WTYTSYao6nKrtAF3X9qTysWzzVeLWK9mFOE31s5vwqBWC
m77jW5UQQtWnbKRhblqi804n9ObWhg9OP4dB9hZgEbR8msExgcLrwyHXJjBn8sod
WALx/LRIV9bWtM9acC402ctJeKYI8S4u33igWbt/Rf6Vi2+Wr+UYeXSufZk8GbFl
HknKcsvKzvwH7yCNqeq5vvgusMZY1vuXe0NS2PESUeVNWGbynnEqk0jvp/Y5Vp3w
n2LtCzDqBzr9GII7x4+sqUCQsLIEXBb1u8kqc7kR2lblXLfxwwoDf6aRWOa2aeaj
QQIDAQABo1MwUTAdBgNVHQ4EFgQUPJwRVXV/bfTCqo+jKF7IfjJW4F8wHwYDVR0j
BBgwFoAUPJwRVXV/bfTCqo+jKF7IfjJW4F8wDwYDVR0TAQH/BAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAgEAY0ZOlWxryKrsB8/ITp0dpuyUBPebTgYMKdywJ1jm9P/A
7RTn/Xg1PhJwHes9ZDLLhIzWmlX79F44WJmY8hLAwavQ3TPSbhdt1qgibhIEg5uq
3cGqvLmpaPo4gTG+iCIj1CYkrgLdQAIxPq53ZDnFv/A25xyxtvwPxMmyRRCyPYoz
rrioGlyixBGGFRmhFlLS72uVxko9MbK9RMsfVzER0SsBQL0SbVKMKTOOzC8NVIYP
268wYlb9VMVXdWPFjyyroNwzuwoztpZB9yHoQrAGDKw5I/k3fBV2V864ByuId0HR
nAfWSEJXBrMC5f+v3BQ2mo7Mc1ALRYHypUB+omLNlnfkMZxvVzQQR+P9aJKlTWhQ
xwMp8nXPJsh6wyZKP65LG2YGgty5mZzoTD+B4PAwToRLrJxeQlO+iNXyGftElDdM
BvFOW/IkaD3o6jA8sIHlgMEOfh1xNIlwofR+TJu/+odqvLzoAx7DkvAqDZN3WbMh
jweC0c3uYWtuu0buFojmPh4yE91uL2OHoS8DuMoRUN6BvvOufnpCtRJymdzUXjRt
k9nv1gdAezi81UF0GMMG9HF75NLZlNxOaqWt+7YoiIMpM/Xafp6I0/KXl2Bl0wuW
31hSk9QWjII7ki7aNnOBDtPa/yrI+AhjlO0hK/sYOGFmqupq//nyKz8RR5fhMA8=
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
	ProjectId      types.String `tfsdk:"project_id"`
	Namespace      types.String `tfsdk:"namespace"`
	Suffix         types.String `tfsdk:"suffix"`
	TLSCertificate types.String `tfsdk:"tls_certificate"`
	SecretId       types.String `tfsdk:"secret_id"`
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
			"tls_certificate": schema.StringAttribute{
				MarkdownDescription: "Add value if not using Let's Encrypt auto renewal certificates",
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
	var cert []byte
	if strings.TrimSpace(data.TLSCertificate.ValueString()) == "" {
		cert = []byte(expiredTLSCert)
	} else {
		cert = []byte(data.TLSCertificate.ValueString())
	}
	if err := addSecretVersion(ctx, t.client, data.ProjectId.ValueString(), secretId, cert); err != nil {
		resp.Diagnostics.AddError("Failed to create tls certificate secret", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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
	var cert []byte
	if strings.TrimSpace(data.TLSCertificate.ValueString()) == "" {
		cert = []byte(expiredTLSCert)
	} else {
		cert = []byte(data.TLSCertificate.ValueString())
	}
	if err := addSecretVersion(ctx, t.client, data.ProjectId.ValueString(), secretId, cert); err != nil {
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
