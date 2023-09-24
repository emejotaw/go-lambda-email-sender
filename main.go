package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/emejotaw/go-lambda-email-sender/client"
	"github.com/emejotaw/go-lambda-email-sender/handler"
)

func main() {

	secretsManagerClient, err := client.NewSecretsManagerClient()

	if err != nil {
		log.Printf("could not get secrets manager client, error: %v", err)
		panic(err)
	}

	handler := handler.NewHandler(secretsManagerClient)
	lambda.Start(handler.HandleFunc)
}
