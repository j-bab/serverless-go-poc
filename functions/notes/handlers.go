package main

import (
	"encoding/json"
	"time"
)

func StringifyItem(item Item) string {
	jsonItem, _ := json.Marshal(item)
	return string(jsonItem)
}

func postRequest(UserId string, Body string) (string, error) {

	now := time.Now()
	requestItem := Item{
		UserId:    UserId,
		Timestamp: now.Unix(),
		Body:      Body,
	}

	item, err := CreateNote(requestItem)

	return StringifyItem(item), err

}

func putRequest(UserId string, Timestamp int64, Body string) (string, error) {

	requestItem := Item{
		UserId:    UserId,
		Timestamp: Timestamp,
		Body:      Body,
	}
	item, err := UpdateNote(requestItem)

	return StringifyItem(item), err
}

func deleteRequest(UserId string, Timestamp int64) (string, error) {

	err := DeleteNote(UserId, Timestamp)
	if err != nil {
		return "Failure", err
	}
	return "Success", nil
}

func getRequest(UserId string, Timestamp int64) (string, error) {
	item, err := GetNote(UserId, Timestamp)

	return StringifyItem(item), err
}

func listRequest(UserId string) (string, error) {
	items, err := ListNotes(UserId)
	stringItems := "["
	for i := 0; i < len(items); i++ {
		stringItems += StringifyItem(items[i])
		if i != len(items)-1 {
			stringItems += ",\n"
		}
	}
	stringItems += "]"
	return stringItems, err
}
