# aws-lambda
> An exprimental implementation of AWS lambda based in the [AWS Lambda for Go](https://github.com/aws/aws-lambda-go) project.

## General info
This project is a proof-of-concept implementation to see how to write, deploy and run Go services on AWS.

## Technologies
* Go - version 1.15

## Setup, Run, Deply
Main reference is [AWS Lambda for Go](https://github.com/aws/aws-lambda-go)

Main steps are build:
```bash
go build -o aws-lambda-go .
or
GOOS=linux go build -o aws-lambda-go .
```

...then, package (macos):

```bash
zip aws-lambda-go.zip ./aws-lambda-go 
```

...then deploy. See [deploying-lambda-apps](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html)

deploying with AWS CLI:
```bash
# requires role 'lambda-url-role' to be created before hand
#   Trusted entity – AWS Lambda.
#   Permissions – AWSLambdaBasicExecutionRole.
#   Role name – lambda-url-role.
# --handler must be the name of the executable in the zip file
aws lambda create-function \
    --function-name aws-lambda-go-cli \
    --runtime go1.x \
    --zip-file fileb://aws-lambda-go.zip \
    --handler aws-lambda-go \
    --role arn:aws:iam::312970875324:role/lambda-url-role
DELETE:
aws lambda delete-function --function-name aws-lambda-go-cli

aws lambda add-permission \
    --function-name aws-lambda-go-cli \
    --action lambda:InvokeFunctionUrl \
    --principal "*" \
    --function-url-auth-type "NONE" \
    --statement-id url

aws lambda create-function-url-config \
    --function-name aws-lambda-go-cli \
    --auth-type NONE
DELETE:
aws lambda delete-function-url-config --function-name aws-lambda-go-cli

aws lambda invoke   \
    --function-name aws-lambda-go-cli   \
    --cli-binary-format raw-in-base64-out  \
    --payload '{"name": "aaa", "age": 20}' output.txt

curl 'https://ko37bhwqbjckm7ia2tqwlndhyy0jhlgt.lambda-url.us-east-1.on.aws/' \
-H 'Content-Type: application/json' \
-d '{"name": "abc", "age": 10}'
```
## Features
Current features include:
* Take name and age and wite it out again

Future improvements:
* Implement CI/CD proses

## Status
Project is: _in progress_

## Contact
Created by [pbreedt](mailto:petrus.breedt@gmail.com)