package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	items, err := ListNotes(request.PathParameters["userId"])
	if err != nil {
		fmt.Println("Failed fetching notes")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "Error", StatusCode: 500}, nil
	}

	// Log and return result
	stringItems := "["
	for i := 0; i < len(items); i++ {
		jsonItem, _ := json.Marshal(items[i])
		stringItems += string(jsonItem)
		if i != len(items)-1 {
			stringItems += ",\n"
		}
	}
	stringItems += "]\n"
	fmt.Println("Found items: ", stringItems)
	return events.APIGatewayProxyResponse{Body: stringItems, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
