package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/emejotaw/go-lambda-email-sender/handler"
)

func main() {

	handler := handler.LambdaHandler{}
	lambda.Start(handler.HandleFunc)
}
