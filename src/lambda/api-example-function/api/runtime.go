package api

import (
	"bytes"
	"context"

	"github.com/michimani/aws-lambda-api-go/alago"
	"github.com/michimani/aws-lambda-api-go/runtime/invocation"
)

func RuntimeAPIInvocationNext(c alago.AlagoClient) (*invocation.NextOutput, error) {
	return invocation.InvocationNext(context.Background(), c)
}

func RuntimeAPIInvocationResponse(c alago.AlagoClient, reqID string, b []byte) (*invocation.ResponseOutput, error) {
	in := invocation.ResponseInput{
		AWSRequestID: reqID,
		Response:     bytes.NewReader(b),
	}

	return invocation.InvocationResponse(context.Background(), c, &in)
}
