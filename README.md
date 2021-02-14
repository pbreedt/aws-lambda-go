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
go build
```

...then, package (macos):

```bash
zip aws-lambda-go.zip ./aws-lambda-go 
```

...then deploy. See [deploying-lambda-apps](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html)

## Features
Current features include:
* Take name and age and wite it out again

Future improvements:
* Implement CI/CD proses

## Status
Project is: _in progress_

## Contact
Created by [pbreedt](mailto:petrus.breedt@gmail.com)