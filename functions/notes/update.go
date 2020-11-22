package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var requestBody ItemRequestBody
	json.Unmarshal([]byte(request.Body), &requestBody)

	timestamp, err := stringToTimestamp(request.PathParameters["timestamp"])

	requestItem := Item{
		UserId:    request.PathParameters["userId"],
		Timestamp: timestamp,
		Body:      requestBody.Body,
	}

	item, err := UpdateNote(requestItem)
	if err != nil {
		fmt.Println("Error updating")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	// Log and return result
	jsonItem, _ := json.Marshal(item)
	stringItem := string(jsonItem) + "\n"
	fmt.Println("Updated item: ", stringItem)
	return events.APIGatewayProxyResponse{Body: stringItem, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
