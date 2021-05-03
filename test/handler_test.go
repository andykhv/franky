package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

func TestGetUserHandler(test *testing.T) {
	request, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		test.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(franky.GetUser)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	trimmedBody := strings.TrimSpace(responseRecorder.Body.String())

	if trimmedBody != UserExpected {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseRecorder.Body.String(), UserExpected)
	}
}

func TestGetRecordsHandler(test *testing.T) {
	request, err := http.NewRequest("GET", "/users/123/records", nil)
	if err != nil {
		test.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(franky.GetRecords)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	trimmedBody := strings.TrimSpace(responseRecorder.Body.String())

	if trimmedBody != RecordsExpected {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseRecorder.Body.String(), RecordsExpected)
	}
}
