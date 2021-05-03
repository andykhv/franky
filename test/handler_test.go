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

	handler := http.HandlerFunc(franky.GetUser)

	testHandler(request, handler, http.StatusOK, UserExpected, test)
}

func TestPostUserHandler(test *testing.T) {
	request, err := http.NewRequest("POST", "/users/123", nil)
	if err != nil {
		test.Fatal(err)
	}

	handler := http.HandlerFunc(franky.PostUser)

	testHandler(request, handler, http.StatusOK, "", test)
}

func TestDeleteUserHandler(test *testing.T) {
	request, err := http.NewRequest("DELETE", "/users/123", nil)
	if err != nil {
		test.Fatal(err)
	}

	handler := http.HandlerFunc(franky.DeleteUser)

	testHandler(request, handler, http.StatusOK, "", test)

}

func TestGetRecordsHandler(test *testing.T) {
	request, err := http.NewRequest("GET", "/users/123/records", nil)
	if err != nil {
		test.Fatal(err)
	}

	handler := http.HandlerFunc(franky.GetRecords)

	testHandler(request, handler, http.StatusOK, RecordsExpected, test)
}

func TestPostRecordHandler(test *testing.T) {
	request, err := http.NewRequest("POST", "/users/123/records", nil)
	if err != nil {
		test.Fatal(err)
	}

	handler := http.HandlerFunc(franky.PostRecord)

	testHandler(request, handler, http.StatusOK, "", test)

}

func testHandler(request *http.Request, handler http.HandlerFunc, expectedStatus int, expectedBody string, test *testing.T) {
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != expectedStatus {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if responseBody := strings.TrimSpace(responseRecorder.Body.String()); responseBody != expectedBody {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseBody, expectedBody)
	}
}
