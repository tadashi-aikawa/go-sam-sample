package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := request.QueryStringParameters["message"]
	if message == "" {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintln("Not found message query"),
			StatusCode: 400,
		}, nil
	}

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintln("Not set SLACK_WEBHOOK_URL"),
			StatusCode: 500,
		}, nil
	}

	s := NewSlack(webhookURL, message, "seigo100%", ":hyakutake_satori:", "", "#times_tadashi-aikawa")
	s.Send()

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintln("Success send"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
