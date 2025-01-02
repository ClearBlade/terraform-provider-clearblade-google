package provider

import (
	"context"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func createSecret(ctx context.Context, client *secretmanager.Client, projectId, secretId string) error {
	parent := "projects/" + projectId
	createReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   parent,
		SecretId: secretId,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	if _, err := client.CreateSecret(ctx, createReq); err != nil {
		return err
	}
	return nil
}

func addSecretVersion(ctx context.Context, client *secretmanager.Client, projectId, secretId string, data []byte) error {
	resource := getSecretResourceName(projectId, secretId)
	addReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: resource,
		Payload: &secretmanagerpb.SecretPayload{
			Data: data,
		},
	}
	if _, err := client.AddSecretVersion(ctx, addReq); err != nil {
		return err
	}
	return nil
}

func getSecretResourceName(projectId, secretId string) string {
	return "projects/" + projectId + "/secrets/" + secretId
}

func getSecretId(namespace, suffix string) string {
	return namespace + suffix
}
