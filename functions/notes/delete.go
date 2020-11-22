package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	timestamp, err := stringToTimestamp(request.PathParameters["timestamp"])

	err = DeleteNote(request.PathParameters["userId"], timestamp)
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	// Log and return result
	fmt.Println("Deleted item: ")

	return events.APIGatewayProxyResponse{Body: "DeletedNote", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
