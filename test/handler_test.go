package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

var userExpected = `{"Id":"id","Email":"email","Password":"password","ApiKey":"apiKey","CreationDate":"creationDate"}`

func TestGetUserHandler(test *testing.T) {
	request, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		test.Fatal(err)
	}

	responseRecoder := httptest.NewRecorder()
	handler := http.HandlerFunc(franky.GetUser)

	handler.ServeHTTP(responseRecoder, request)

	if status := responseRecoder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if responseRecoder.Body.String() != userExpected {
		test.Errorf("handler returned unexpected body: got %v want %v",
			responseRecoder.Body.String(), userExpected)
	}
}
