package provider

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/big"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RandomStringResource{}
var _ resource.ResourceWithImportState = &RandomStringResource{}

type randomStringType string

const (
	password        = randomStringType("password")
	registrationKey = randomStringType("registration_key")
)

func NewRandomStringResource() resource.Resource {
	return &RandomStringResource{}
}

// RandomStringResource defines the resource implementation.
type RandomStringResource struct {
	client *secretmanager.Client
}

// RandomStringResourceModel describes the resource data model.
type RandomStringResourceModel struct {
	ProjectId types.String `tfsdk:"project_id"`
	Namespace types.String `tfsdk:"namespace"`
	Suffix    types.String `tfsdk:"suffix"`
	Type      types.String `tfsdk:"type"`
	Length    types.Int32  `tfsdk:"length"`
	SecretId  types.String `tfsdk:"secret_id"`
	Value     types.String `tfsdk:"value"`
}

func (r *RandomStringResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_random_string"
}

func (r *RandomStringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Random string resource to create a random password and/or ClearBlade registration key",

		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				MarkdownDescription: "GCP project Id",
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
			"type": schema.StringAttribute{
				MarkdownDescription: "Random string type",
				Required:            true,
			},
			"length": schema.Int32Attribute{
				MarkdownDescription: "Length of random string",
				Required:            true,
			},
			"secret_id": schema.StringAttribute{
				Computed: true,
			},
			"value": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *RandomStringResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (r *RandomStringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RandomStringResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch randomStringType(data.Type.ValueString()) {
	case password:
		if err := r.createPassword(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Failed to create password", err.Error())
			return
		}
	case registrationKey:
		if err := r.createRegistrationKey(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Failed to create registration key", err.Error())
			return
		}
	default:
		resp.Diagnostics.AddError("Invalid type attribute", data.Type.ValueString())
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RandomStringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RandomStringResourceModel

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
	rval, err := r.client.AccessSecretVersion(ctx, secReq)
	if err != nil {
		resp.Diagnostics.AddError("Faled to get MEK secret", err.Error())
		return
	}
	data.SecretId = types.StringValue(secretId)
	switch randomStringType(data.Type.ValueString()) {
	case password:
		data.Value = types.StringValue(hashPassword(string(rval.Payload.Data)))
	default:
		data.Value = types.StringValue(string(rval.Payload.Data))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RandomStringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RandomStringResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch randomStringType(data.Type.ValueString()) {
	case password:
		if err := r.updatePassword(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Failed to update password", err.Error())
			return
		}
	case registrationKey:
		if err := r.updateRegistrationKey(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Failed to update registration key", err.Error())
			return
		}
	default:
		resp.Diagnostics.AddError("Invalid type attribute", data.Type.ValueString())
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RandomStringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RandomStringResourceModel

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
	if err := r.client.DeleteSecret(ctx, delReq); err != nil {
		resp.Diagnostics.AddError("Failed to delete password", err.Error())
		return
	}
}

func (r *RandomStringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RandomStringResource) createPassword(ctx context.Context, data *RandomStringResourceModel) error {
	if err := validatePasswordLength(data.Length.ValueInt32()); err != nil {
		return fmt.Errorf("Invalid password length: %w", err)
	}
	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	if err := createSecret(ctx, r.client, data.ProjectId.ValueString(), secretId); err != nil {
		return fmt.Errorf("Failed to create secret: %w", err)
	}
	password, err := generateRandomString(int(data.Length.ValueInt32()))
	if err != nil {
		return fmt.Errorf("Failed to generate random password: %w", err)
	}
	if err := addSecretVersion(ctx, r.client, data.ProjectId.ValueString(), secretId, []byte(password)); err != nil {
		return fmt.Errorf("Failed to add password to secret: %w", err)
	}
	data.Value = types.StringValue(hashPassword(password))
	data.SecretId = types.StringValue(secretId)
	return nil
}

func (r *RandomStringResource) updatePassword(ctx context.Context, data *RandomStringResourceModel) error {
	if err := validatePasswordLength(data.Length.ValueInt32()); err != nil {
		return fmt.Errorf("Invalid password length: %w", err)
	}
	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	data.SecretId = types.StringValue(secretId)
	password, err := generateRandomString(int(data.Length.ValueInt32()))
	if err != nil {
		return fmt.Errorf("Failed to generate random password: %w", err)
	}
	if err := addSecretVersion(ctx, r.client, data.ProjectId.ValueString(), secretId, []byte(password)); err != nil {
		return fmt.Errorf("Failed to add password to secret: %w", err)
	}
	data.Value = types.StringValue(hashPassword(password))
	return nil
}

func (r *RandomStringResource) createRegistrationKey(ctx context.Context, data *RandomStringResourceModel) error {
	if err := validateRegistrationKeyLength(data.Length.ValueInt32()); err != nil {
		return fmt.Errorf("Invalid registration key length: %w", err)
	}
	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	if err := createSecret(ctx, r.client, data.ProjectId.ValueString(), secretId); err != nil {
		return fmt.Errorf("Failed to create secret: %w", err)
	}
	registrationKey, err := generateRandomString(int(data.Length.ValueInt32()))
	if err != nil {
		return fmt.Errorf("Failed to generate random registration key: %w", err)
	}
	if err := addSecretVersion(ctx, r.client, data.ProjectId.ValueString(), secretId, []byte(registrationKey)); err != nil {
		return fmt.Errorf("Failed to add registration key to secret: %w", err)
	}
	data.Value = types.StringValue(registrationKey)
	data.SecretId = types.StringValue(secretId)
	return nil
}

func (r *RandomStringResource) updateRegistrationKey(ctx context.Context, data *RandomStringResourceModel) error {
	if err := validateRegistrationKeyLength(data.Length.ValueInt32()); err != nil {
		return fmt.Errorf("Invalid registration key length: %w", err)
	}
	secretId := getSecretId(data.Namespace.ValueString(), data.Suffix.ValueString())
	data.SecretId = types.StringValue(secretId)
	registrationKey, err := generateRandomString(int(data.Length.ValueInt32()))
	if err != nil {
		return fmt.Errorf("Failed to generate random registration key: %w", err)
	}
	if err := addSecretVersion(ctx, r.client, data.ProjectId.ValueString(), secretId, []byte(registrationKey)); err != nil {
		return fmt.Errorf("Failed to add registration key to secret: %w", err)
	}
	data.Value = types.StringValue(registrationKey)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func hashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func validatePasswordLength(length int32) error {
	if length < 6 {
		return fmt.Errorf("Password length must be greater than or equal to 6")
	} else if length > 30 {
		return fmt.Errorf("Password length cannot be greater than 30")
	} else {
		return nil
	}
}

func validateRegistrationKeyLength(length int32) error {
	if length < 6 {
		return fmt.Errorf("Registration key length must be greater than or equal to 6")
	} else if length > 10 {
		return fmt.Errorf("Registration key length cannot be greater than 10")
	} else {
		return nil
	}
}
