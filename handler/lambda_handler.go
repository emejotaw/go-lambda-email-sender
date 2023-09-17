package handler

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/emejotaw/go-lambda-email-sender/client"
	"github.com/emejotaw/go-lambda-email-sender/encoder"
	"github.com/emejotaw/go-lambda-email-sender/types"
)

type lambdaHandler struct {
	secretsManagerClient *client.SecretsManagerClient
}

func NewHandler(secretsManagerClient *client.SecretsManagerClient) *lambdaHandler {
	return &lambdaHandler{secretsManagerClient: secretsManagerClient}
}

func (lh *lambdaHandler) HandleFunc(ctx context.Context, sqsEvent events.SQSEvent) error {

	encoder := encoder.New()

	log.Printf("Event received: %v", sqsEvent)

	for _, record := range sqsEvent.Records {

		requestDTO := &types.EmailRequestDTO{}
		err := encoder.Encode(record.Body, requestDTO)

		if err != nil {
			log.Fatalf("could not parse json record, error: %v", err)
			return err
		}

		//lh.SendEmail(requestDTO)
		log.Printf("request parsed successfully, data: %v", requestDTO)
	}

	return nil
}

func (lh *lambdaHandler) SendEmail(requestDTO *types.EmailRequestDTO) {

	secretID := os.Getenv("EMAIL_SECRET_ID")
	secretsManagerClient := lh.secretsManagerClient
	secretString, err := secretsManagerClient.Get(secretID)

	if err != nil {
		log.Printf("could not get secret, error: %v", err.Error())
		//TODO
	}

	log.Printf("Get secret %v successfully with data %v", secretID, secretString)
	//Add code to send the email
}
