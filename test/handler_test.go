package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	franky "github.com/andykhv/franky/pkg"
	"github.com/gorilla/mux"
)

var (
	dao     = franky.NewTestDAO()
	handler = franky.NewFrankyHandler(&dao)
)

const (
	usersRoute   = "/users/{id:[0-9]+}"
	recordsRoute = "/users/{id:[0-9]+}/records"
)

func TestGetUserHandler(test *testing.T) {
	const path = "/users/123"

	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		test.Fatal(err)
	}

	testHandler(request, handler.GetUser, usersRoute, http.StatusOK, UserExpected, test)
}

func TestPostUserHandler(test *testing.T) {
	const path = "/users/123"

	request, err := http.NewRequest("POST", path, bytes.NewBufferString(UserExpected))
	if err != nil {
		test.Fatal(err)
	}

	testHandler(request, handler.PostUser, usersRoute, http.StatusOK, "", test)
}

func TestDeleteUserHandler(test *testing.T) {
	const path = "/users/123"

	request, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		test.Fatal(err)
	}

	testHandler(request, handler.DeleteUser, usersRoute, http.StatusOK, "", test)

}

func TestGetRecordsHandler(test *testing.T) {
	const path = "/users/123/records"

	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		test.Fatal(err)
	}

	testHandler(request, handler.GetRecords, recordsRoute, http.StatusOK, RecordsExpected, test)
}

func TestPostRecordHandler(test *testing.T) {
	const path = "/users/123/records"

	request, err := http.NewRequest("POST", path, bytes.NewBufferString(RecordsExpected))
	if err != nil {
		test.Fatal(err)
	}

	testHandler(request, handler.PostRecord, recordsRoute, http.StatusOK, "", test)
}

func testHandler(request *http.Request, handler http.HandlerFunc, path string, expectedStatus int, expectedBody string, test *testing.T) {
	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, handler)
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != expectedStatus {
		test.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if responseBody := strings.TrimSpace(responseRecorder.Body.String()); responseBody != expectedBody {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseBody, expectedBody)
	}
}
