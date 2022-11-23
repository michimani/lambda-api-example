package main

import (
	"lambda-api-example/resource"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type LambdaApiExampleStackProps struct {
	awscdk.StackProps
}

func NewLambdaApiExampleStack(scope constructs.Construct, id string, props *LambdaApiExampleStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	_, err := resource.LambdaFunctionGo(stack, &resource.LambdaFunctionVariables{
		Name:      "api-example-function",
		AssetPath: "./src/lambda/api-example-function/bin",
		Memory:    128,
		Timeout:   15,
	})

	if err != nil {
		panic(err)
	}

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaApiExampleStack(app, "LambdaApiExampleStack", &LambdaApiExampleStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
