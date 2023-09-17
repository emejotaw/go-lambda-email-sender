package client

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type secretsManagerClient struct {
	client *secretsmanager.Client
}

func NewSecretsManagerClient() (*secretsManagerClient, error) {

	client, err := newSecretsManager()
	return &secretsManagerClient{client: client}, err
}

func newSecretsManager() (*secretsmanager.Client, error) {

	ctx := context.Background()
	endpointResolver := aws.EndpointResolverWithOptionsFunc(withEndpointResolver)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithEndpointResolverWithOptions(endpointResolver))

	if err != nil {
		log.Printf("configuration error: %v", err.Error())
		return nil, err
	}

	return secretsmanager.NewFromConfig(cfg), nil
}

func withEndpointResolver(service, region string, options ...interface{}) (aws.Endpoint, error) {

	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")

	if awsEndpoint != "" {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           awsEndpoint,
			SigningRegion: awsRegion,
		}, nil
	}

	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
}

func (smc *secretsManagerClient) Get(secretID string) (string, error) {

	ctx := context.TODO()
	client := smc.client
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	output, err := client.GetSecretValue(ctx, input)

	if err != nil {
		return "", err
	}

	return aws.ToString(output.SecretString), nil
}
