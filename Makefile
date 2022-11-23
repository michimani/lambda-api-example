.PHONY: build-func synth prepare

build-func:
	cd src/lambda/api-example-function && GOARCH=amd64 GOOS=linux go build -o bin/main

synth:
	cdk synth

prepare: build-func synth
