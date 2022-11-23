package main

import (
	"context"
	"hello/api"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/michimani/aws-lambda-api-go/alago"
	"github.com/rs/zerolog"
)

type Response struct {
	AWSRequestID string `json:"awsRequestId"`
}

func handleRequest(ctx context.Context) (*Response, error) {
	ac, err := alago.NewClient(&alago.NewClientInput{})
	if err != nil {
		return nil, err
	}

	// Get AwsRequestID from Runtime API
	res, err := api.RuntimeAPIInvocationNext(ac)
	if err != nil {
		return nil, err
	}
	logger := zerolog.New(os.Stdout).With().Str("awsRequestId", res.AWSRequestID).Caller().Timestamp().Logger()

	logger.Info().Msg("start handler")
	defer logger.Info().Msg("finish handler")

	return &Response{
		AWSRequestID: res.AWSRequestID,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
