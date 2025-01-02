package provider

import (
	"bytes"
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MEKResource{}
var _ resource.ResourceWithImportState = &MEKResource{}

func NewMEKResource() resource.Resource {
	return &MEKResource{}
}

// MEKResource defines the resource implementation.
type MEKResource struct {
	client *secretmanager.Client
}

// MEKResourceModel describes the resource data model.
type MEKResourceModel struct {
	ProjectId types.String `tfsdk:"project_id"`
	Namespace types.String `tfsdk:"namespace"`
	Suffix    types.String `tfsdk:"suffix"`
	SecretId  types.String `tfsdk:"secret_id"`
	Key       types.String `tfsdk:"key"`
}

func (m *MEKResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mek"
}

func (m *MEKResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "ClearBlade Master Encryption Key",

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
			"secret_id": schema.StringAttribute{
				Computed: true,
			},
			"key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (m *MEKResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	m.client = client
}

func (m *MEKResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MEKResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	if err := createSecret(ctx, m.client, data.ProjectId.ValueString(), secretId); err != nil {
		resp.Diagnostics.AddError("Failed to create secret", err.Error())
		return
	}
	data.SecretId = types.StringValue(secretId)
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		resp.Diagnostics.AddError("Failed to create new MEK", err.Error())
		return
	}
	mrw := &mekGCPReaderWriter{
		ctx:       ctx,
		client:    m.client,
		projectId: data.ProjectId.ValueString(),
		secretId:  secretId,
	}
	w := keyset.NewJSONWriter(mrw)
	if err := insecurecleartextkeyset.Write(kh, w); err != nil {
		resp.Diagnostics.AddError("Failed to write MEK to GCP Secrets", err.Error())
		return
	}

	data.Key = types.StringValue(kh.String())
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created and stored MEK to GCP secrets")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (m *MEKResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MEKResourceModel

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
	rval, err := m.client.AccessSecretVersion(ctx, secReq)
	if err != nil {
		resp.Diagnostics.AddError("Faled to get MEK secret", err.Error())
		return
	}
	kh, err := insecurecleartextkeyset.Read(keyset.NewJSONReader(bytes.NewReader(rval.Payload.Data)))
	if err != nil {
		resp.Diagnostics.AddError("Failed to read MEK", err.Error())
		return
	}

	data.Key = types.StringValue(kh.String())
	data.SecretId = types.StringValue(secretId)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (m *MEKResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MEKResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		resp.Diagnostics.AddError("Failed to create new MEK", err.Error())
		return
	}
	mrw := &mekGCPReaderWriter{
		ctx:       ctx,
		client:    m.client,
		projectId: data.ProjectId.ValueString(),
		secretId:  secretId,
	}
	w := keyset.NewJSONWriter(mrw)
	if err := insecurecleartextkeyset.Write(kh, w); err != nil {
		resp.Diagnostics.AddError("Failed to write MEK to GCP Secrets", err.Error())
		return
	}

	data.Key = types.StringValue(kh.String())
	data.SecretId = types.StringValue(secretId)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (m *MEKResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MEKResourceModel

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
	if err := m.client.DeleteSecret(ctx, delReq); err != nil {
		resp.Diagnostics.AddError("Failed to delete MEK", err.Error())
		return
	}
}

func (m *MEKResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type mekGCPReaderWriter struct {
	ctx       context.Context
	client    *secretmanager.Client
	projectId string
	secretId  string
}

func (m *mekGCPReaderWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m *mekGCPReaderWriter) Write(p []byte) (n int, err error) {
	if err := addSecretVersion(m.ctx, m.client, m.projectId, m.secretId, p); err != nil {
		return -1, err
	}
	return len(p), nil
}
