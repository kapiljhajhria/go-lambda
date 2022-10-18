package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kapiljhajhria/go-lambda/handlers"
)

func main() {
	lambda.Start(handlers.HelloWorldHandler)
}
