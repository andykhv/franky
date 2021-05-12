package test

import (
	"fmt"
	"net/http"
	"testing"
)

var getRecordsTests = []struct {
	userId         string
	expectedStatus int
	expectedBody   string
}{
	{"123", http.StatusOK, records},
}

func TestGetRecordsHandler(tester *testing.T) {
	for _, t := range getRecordsTests {
		path := fmt.Sprintf("/users/%s/records", t.userId)
		request, err := http.NewRequest("GET", path, nil)

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.GetRecords, recordsRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
