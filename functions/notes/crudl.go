package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"os"
	"strconv"
)

type ItemRequestBody struct {
	Body string `json:"body"`
}

type Item struct {
	UserId    string `json:"userId"`
	Timestamp int64  `json:"timestamp"`
	Body      string `json:"body"`
}

func getDynamoDb() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess)
}

func getTableName() *string {
	return aws.String(os.Getenv("notesTable"))
}

func timestampToString(Timestamp int64) string {
	return strconv.FormatInt(Timestamp, 10)
}
func stringToTimestamp(TimestampString string) (int64, error) {
	return strconv.ParseInt(TimestampString, 10, 64)
}

func CreateNote(item Item) (Item, error) {

	svc := getDynamoDb()
	//Map item to a format dynamoDb can use
	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println("Error marshalling map:")
		fmt.Println(err.Error())
		return item, err
	}
	// Create item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: getTableName(),
	}
	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Error calling PutItem:")
		fmt.Println(err.Error())
		return item, err
	}
	return item, nil
}

func GetNote(UserId string, Timestamp int64) (Item, error) {
	svc := getDynamoDb()
	item := Item{}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: getTableName(),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(UserId),
			},
			"timestamp": {
				N: aws.String(timestampToString(Timestamp)),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}
	// Unmarshall the result in to an Item
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}
	return item, nil
}

func UpdateNote(item Item) (Item, error) {
	svc := getDynamoDb()

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":b": {
				S: aws.String(item.Body),
			},
		},
		TableName: getTableName(),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(item.UserId),
			},
			"timestamp": {
				N: aws.String(timestampToString(item.Timestamp)),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set body = :b"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return item, err
	}
	return item, nil
}

func DeleteNote(UserId string, Timestamp int64) error {
	svc := getDynamoDb()
	// Perform the delete
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(UserId),
			},
			"timestamp": {
				N: aws.String(timestampToString(Timestamp)),
			},
		},
		TableName: getTableName(),
	}
	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ListNotes(UserId string) ([]Item, error) {
	svc := getDynamoDb()
	items := []Item{}
	//filter by user id
	userIdFilter := expression.Name("userId").Equal(expression.Value(UserId))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("userId"), expression.Name("timestamp"), expression.Name("body"))

	expr, err := expression.NewBuilder().WithFilter(userIdFilter).WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return items, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 getTableName(),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	fmt.Println("Result", result)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		return items, err
	}

	numItems := 0
	for _, i := range result.Items {
		item := Item{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return items, err
		}
		items = append(items, item)
		numItems++
	}

	fmt.Println("Found", numItems, " notes ")
	if err != nil {
		fmt.Println(err.Error())
		return items, err
	}

	return items, nil
}
