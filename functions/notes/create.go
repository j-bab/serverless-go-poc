package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var requestBody ItemRequestBody
	json.Unmarshal([]byte(request.Body), &requestBody)

	now := time.Now()
	requestItem := Item{
		UserId:    request.PathParameters["userId"],
		Timestamp: now.Unix(),
		Body:      requestBody.Body,
	}

	item, err := CreateNote(requestItem)
	if err != nil {
		fmt.Println("Error creating item")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	// Log and return result
	jsonItem, _ := json.Marshal(item)
	stringItem := string(jsonItem) + "\n"
	fmt.Println("Wrote item: ", stringItem)
	return events.APIGatewayProxyResponse{Body: stringItem, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
