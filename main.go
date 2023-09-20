package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// event is actually of type events.LambdaFunctionURLRequest
type AmzEvent struct {
	Body string `json:"body"`
}

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, evt AmzEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", evt.Body), nil
}

func main() {
	lambda.Start(HandleRequest)
}
