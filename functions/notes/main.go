package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	Timestamp, err := stringToTimestamp(request.PathParameters["timestamp"])
	UserId := request.PathParameters["userId"]
	var responseBody = ""

	fmt.Println("--" + request.HTTPMethod + "Request. UserId:" + UserId + " timestamp:" + request.PathParameters["timestamp"])

	switch request.HTTPMethod {
	case "GET":
		if Timestamp > 0 {
			responseBody, err = getRequest(UserId, Timestamp)
		} else {
			responseBody, err = listRequest(UserId)
		}
	case "PUT", "POST":
		var requestBody ItemRequestBody
		err = json.Unmarshal([]byte(request.Body), &requestBody)
		if Timestamp > 0 {
			responseBody, err = putRequest(UserId, Timestamp, requestBody.Body)
		} else {
			responseBody, err = postRequest(UserId, requestBody.Body)
		}
	case "DELETE":
		responseBody, err = deleteRequest(UserId, Timestamp)
	}

	if err != nil {
		fmt.Println("Finished: Error processing note")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	fmt.Println(responseBody)
	fmt.Println("FINISHED")
	return events.APIGatewayProxyResponse{Body: responseBody, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
