package handlers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func HelloWorldHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO implement
	if request.HTTPMethod == "GET" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       fmt.Sprintf("Hello World!"),
		}, nil
	}
	fmt.Println("Hello Work-----------------------------------> called")
	return events.APIGatewayProxyResponse{}, nil
}
