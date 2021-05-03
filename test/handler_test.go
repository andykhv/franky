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

	responseRecoder := httptest.NewRecorder()
	handler := http.HandlerFunc(franky.GetUser)

	handler.ServeHTTP(responseRecoder, request)

	if status := responseRecoder.Code; status != http.StatusOK {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	trimmedBody := strings.TrimSpace(responseRecoder.Body.String())

	if trimmedBody != UserExpected {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseRecoder.Body.String(), UserExpected)
	}
}
