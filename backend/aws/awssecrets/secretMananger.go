package awssecrets

import (
	"context"
	"fmt"

	"github.com/GaryBrownEEngr/go_web_dev/backend/models"
	"github.com/GaryBrownEEngr/go_web_dev/backend/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func NewAwsSecretManager(getenv func(string) string) (models.SecretStore, error) {
	secretName := getenv("AWS_SECRET_NAME")
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
	region := getenv("AWS_REGION")
	// region := "us-west-2"

	// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("Error while loading AWS default config: %w", err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("Error while getting secrets: %w", err)
	}

	if result.SecretString == nil {
		return "", fmt.Errorf("Secret string is missing")
	}

	return *result.SecretString, nil
}
