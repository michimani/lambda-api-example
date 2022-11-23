package api

import (
	"context"

	"github.com/michimani/aws-lambda-api-go/alago"
	"github.com/michimani/aws-lambda-api-go/runtime/invocation"
)

func RuntimeAPIInvocationNext(c alago.AlagoClient) (*invocation.NextOutput, error) {
	return invocation.InvocationNext(context.Background(), c)
}
