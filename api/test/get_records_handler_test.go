package test

import (
	"fmt"
	"net/http"
	"testing"
)

var getRecordsTests = []struct {
	userId         string
	startDate      string
	endDate        string
	expectedStatus int
	expectedBody   string
}{
	{"123", `01 Jul 21 19:02 %2b0000`, `01 Jul 21 19:02 %2b0000`, http.StatusOK, records},
}

func TestGetRecordsHandler(tester *testing.T) {
	for _, t := range getRecordsTests {
		path := fmt.Sprintf("/users/%s/records?startDate=%s&endDate=%s", t.userId, t.startDate, t.endDate)
		request, err := http.NewRequest("GET", path, nil)

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.GetRecords, recordsRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
