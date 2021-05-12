package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

var postRecordTests = []struct {
	userId         string
	records        string
	expectedStatus int
	expectedBody   string
}{
	{"123", records, http.StatusOK, ""},
}

func TestPostRecordHandler(tester *testing.T) {
	for _, t := range postRecordTests {
		path := fmt.Sprintf("/users/%s/records", t.userId)

		request, err := http.NewRequest("POST", path, bytes.NewBufferString(t.records))
		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.PostRecord, recordsRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
