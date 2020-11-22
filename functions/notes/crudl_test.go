package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetTableName(t *testing.T) {
	expectedVal := "TEST_VALUE"
	os.Setenv("notesTable", expectedVal)

	testResult := getTableName()
	if strings.Compare(testResult, expectedVal) != 0 {
		t.Errorf("getTableName was incorrect, got: %s, want: %s.", testResult, expectedVal)
	}
}
func TestTimestampToString(t *testing.T) {

	tables := []struct {
		testVal     int64
		expectedVal string
	}{
		{0, "0"},
		{123, "123"},
		{1, "1"},
		{999, "999"},
	}

	for _, table := range tables {
		testResult := timestampToString(table.testVal)
		if strings.Compare(testResult, table.expectedVal) != 0 {
			t.Errorf("timestampToString was incorrect, got: %s, want: %s.", testResult, table.expectedVal)
		}
	}

}
func TestStringToTimestamp(t *testing.T) {

	tables := []struct {
		testVal     string
		expectedVal int64
	}{
		{"123", 123},
		{"abc", 0},
		{"123abc", 0},
		{"0", 0},
		{"1", 1},
	}

	for _, table := range tables {
		testResult, _ := stringToTimestamp(table.testVal)
		if testResult != table.expectedVal {
			t.Errorf("stringToTimestamp was incorrect, got: %d, want: %d.", testResult, table.expectedVal)
		}
	}

}
