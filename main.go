package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/emejotaw/go-lambda-email-sender/client"
	"github.com/emejotaw/go-lambda-email-sender/handler"
)

func main() {

	secretsManagerClient, err := client.NewSecretsManagerClient()

	if err != nil {
		panic(err)
	}

	handler := handler.NewHandler(secretsManagerClient)
	lambda.Start(handler.HandleFunc)
}
