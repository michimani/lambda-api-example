Lambda API Example
===

This is an example of AWS Lambda Function that calls Runtime API, Extension API, Telemetry API, and Logs API. Resources are built with AWS CDK (Go).

- AWS CDK: `2.51.1`
- Go: `1.19`

# Usage

## Preparing for deploy

Build function code, and generate CloudFormation template.

```bash
make prepare
```

## Deploy

Check diff (optional)

```bash
cdk diff
```

Deploy

```bash
cdk deploy
```

## Invoke function

```bash
aws lambda invoke \
--function-name 'api-example-function' \
--invocation-type 'RequestResponse' \
--cli-binary-format 'raw-in-base64-out' \
--payload '{}' \
--log-type 'Tail' \
/dev/stdout \
| tee >(jq -sr '.[0]') \
| jq -sr '.[1] | .LogResult' | base64 -d
```

# License

[MIT](https://github.com/michimani/lambda-api-example/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)