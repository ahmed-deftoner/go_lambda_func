package main

import "github.com/aws/aws-lambda-go/lambda"

type MyEvent struct {
	Name string `json:"what is ur name?"`
	Age  int    `json:"how old are u?"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
