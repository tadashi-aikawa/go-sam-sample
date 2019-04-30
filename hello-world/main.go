package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	// "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// s := NewSlack(os.Getenv("SLACK_WEBHOOK_URL"), "test", "gopher", ":gopher:", "", "#times_tadashi-aikawa")
	s := NewSlack("", "test", "gopher", ":gopher:", "", "#times_tadashi-aikawa")
	s.Send()

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Success send"),
		StatusCode: 200,
	}, nil
}

func main() {
	//lambda.Start(handler)
	fmt.Sprint("Hello")
}
