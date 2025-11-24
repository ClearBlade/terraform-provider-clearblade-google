package provider

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilehostingHmacResource{}
var _ resource.ResourceWithImportState = &FilehostingHmacResource{}

func NewFilehostingHmacResource() resource.Resource {
	return &FilehostingHmacResource{}
}

// FilehostingHmacResource defines the resource implementation.
type FilehostingHmacResource struct {
	client *secretmanager.Client
}

// FilehostingHmacResourceModel describes the resource data model.
type FilehostingHmacResourceModel struct {
	ProjectId types.String `tfsdk:"project_id"`
	Namespace types.String `tfsdk:"namespace"`
	Suffix    types.String `tfsdk:"suffix"`
	SecretId  types.String `tfsdk:"secret_id"`
	HmacKey   types.String `tfsdk:"hmac_key"`
}

func (f *FilehostingHmacResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filehosting_hmac_secret"
}

func (f *FilehostingHmacResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				MarkdownDescription: "Map of certificate names to PEM encoded certificate strings. If using ACME, this should be an empty map.",
				Required:            true,
			},
			"hmac_key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (t *FilehostingHmacResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (t *FilehostingHmacResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FilehostingHmacResourceModel

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

func (t *FilehostingHmacResourceModel) getSecretBytes() ([]byte, error) {
	if !t.HmacKey.IsNull() && t.HmacKey.ValueString() != "" {
		return []byte(t.HmacKey.ValueString()), nil
	}

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, hex.EncodedLen(len(b)))
	_ = hex.Encode(dst, b)
	return dst, nil
}

func (t *FilehostingHmacResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FilehostingHmacResourceModel

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

func (t *FilehostingHmacResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FilehostingHmacResourceModel

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

func (t *FilehostingHmacResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilehostingHmacResourceModel

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

func (e *FilehostingHmacResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
