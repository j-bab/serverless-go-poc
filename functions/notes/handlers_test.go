package main

import (
	"strings"
	"testing"
)

func TestStringifyItem(t *testing.T) {

	tables := []struct {
		testVal     Item
		expectedVal string
	}{
		{Item{UserId: "12", Timestamp: 123456123456, Body: "this is my note"},
			"{\"userId\":\"12\",\"timestamp\":123456123456,\"body\":\"this is my note\"}"},
		{Item{UserId: "abss", Timestamp: 1234564456, Body: "üäÜÄöÖ!§$%&/()"},
			"{\"userId\":\"abss\",\"timestamp\":1234564456,\"body\":\"üäÜÄöÖ!§$%\\u0026/()\"}"},
	}

	for _, table := range tables {
		testResult := StringifyItem(table.testVal)
		if strings.Compare(testResult, table.expectedVal) != 0 {
			t.Errorf("StringifyItem was incorrect, got: %s, want: %s.", testResult, table.expectedVal)
		}
	}

}
