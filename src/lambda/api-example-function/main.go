package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hello/api"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/michimani/aws-lambda-api-go/alago"
	"github.com/rs/zerolog"
)

type Response struct {
	AWSRequestID string `json:"awsRequestId"`
	Message      string `json:"message"`
}

type apiType int

const (
	apiTypeNone apiType = iota
	apiTypeRuntimeInvocationResponse
)

type Payload struct {
	APIType     apiType `json:"apiType"`
	ReturnError bool    `json:"returnError"`
}

func handleRequest(ctx context.Context, p Payload) (*Response, error) {
	ac, err := alago.NewClient(&alago.NewClientInput{})
	if err != nil {
		return nil, err
	}

	// Get AwsRequestID from Runtime API
	res, err := api.RuntimeAPIInvocationNext(ac)
	if err != nil {
		return nil, err
	}

	reqID := res.AWSRequestID
	logger := zerolog.New(os.Stdout).With().Str("awsRequestId", reqID).Caller().Timestamp().Logger()

	// call Lambda API
	if err := callLambdaAPI(ac, reqID, p.APIType); err != nil {
		logger.Error().Msgf("Failed to call lambda api. err:%v", err)
	}

	logger.Info().Msg("start handler")
	defer logger.Info().Msg("finish handler")

	return &Response{
		AWSRequestID: reqID,
		Message:      "Hello AWS Lambda!",
	}, nil
}

func callLambdaAPI(ac *alago.Client, reqID string, at apiType) error {
	switch at {
	case apiTypeNone:
		return nil
	case apiTypeRuntimeInvocationResponse:
		customRes := Response{
			AWSRequestID: reqID,
			Message:      "This message is set by Runtime API (invocation/response).",
		}
		b, jerr := json.Marshal(customRes)
		if jerr != nil {
			return jerr
		}

		_, err := api.RuntimeAPIInvocationResponse(ac, reqID, b)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown api type: %d", at)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
