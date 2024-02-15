package gcpsecretmanager

import (
	"context"
	"fmt"

	gcpsecretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
)

func NewGcpSecretManager(getenv func(string) string) (models.SecretStore, error) {
	secretName := getenv("GCP_SECRET_NAME")
	// AWS_SECRET_NAME == "AppRunner/GoWebDev"
	secretVal, err := GetSecret(getenv, secretName)
	if err != nil {
		return nil, err
	}

	result, err := utils.NewSecretManager(secretVal)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// https://cloud.google.com/secret-manager/docs/reference/libraries#client-libraries-usage-go
func GetSecret(getenv func(string) string, secretName string) (string, error) {
	// GCP project in which to store secrets in Secret Manager.
	projectID := getenv("GCP_PROJECT_ID")

	// Create the client.
	ctx := context.Background()
	client, err := gcpsecretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Build the request.
	secretPath := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretPath,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", err
	}

	resultString := string(result.Payload.Data)
	return resultString, nil
}
