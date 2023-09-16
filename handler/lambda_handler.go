package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/emejotaw/go-lambda-email-sender/encoder"
	"github.com/emejotaw/go-lambda-email-sender/types"
)

type LambdaHandler struct {
}

func (lh *LambdaHandler) HandleFunc(ctx context.Context, sqsEvent events.SQSEvent) error {

	encoder := encoder.New()

	log.Printf("Event received: %v", sqsEvent)

	for _, record := range sqsEvent.Records {

		requestDTO := &types.EmailRequestDTO{}
		err := encoder.Encode(record.Body, requestDTO)

		if err != nil {
			log.Fatalf("could not parse json record, error: %v", err)
			return err
		}

		log.Printf("request parsed successfully, data: %v", requestDTO)
	}

	return nil
}
