package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	timestamp, err := stringToTimestamp(request.PathParameters["timestamp"])
	item, err := GetNote(request.PathParameters["userId"], timestamp)
	if err != nil {
		fmt.Println("Got error fetching record")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	// Log and return result
	jsonItem, _ := json.Marshal(item)
	stringItem := string(jsonItem) + "\n"
	fmt.Println("Retrieved item: ", stringItem)
	return events.APIGatewayProxyResponse{Body: stringItem, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
