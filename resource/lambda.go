package resource

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaFunctionVariables struct {
	Name        string
	Description string
	AssetPath   string
	Memory      float64
	Timeout     float64
}

func LambdaFunctionGo(stack constructs.Construct, variables *LambdaFunctionVariables) (awslambda.Function, error) {
	if variables == nil {
		return nil, fmt.Errorf("LambdaFunctionVariables is nil")
	}
	if variables.Name == "" {
		return nil, fmt.Errorf("LambdaFunctionVariables.Name is empty")
	}

	return awslambda.NewFunction(stack, jsii.String(variables.Name), &awslambda.FunctionProps{
		FunctionName: jsii.String(variables.Name),
		Description:  jsii.String(variables.Description),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String(variables.AssetPath), nil),
		MemorySize:   &variables.Memory,
		Timeout:      awscdk.Duration_Seconds(&variables.Timeout),
	}), nil
}
