package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/rs/zerolog"
)

func handleRequest(ctx context.Context) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	reqID := lc.AwsRequestID
	logger := zerolog.New(os.Stdout).With().Str("awsRequestId", reqID).Caller().Timestamp().Logger()

	logger.Info().Msg("start handler")
	defer logger.Info().Msg("finish handler")

	return reqID, nil
}

func main() {
	lambda.Start(handleRequest)
}
