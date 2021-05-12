package test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	franky "github.com/andykhv/franky/pkg"
	"github.com/gorilla/mux"
)

const (
	usersRoute   = "/users/{id:[0-9]+}"
	recordsRoute = "/users/{id:[0-9]+}/records"
)

var (
	dao       = NewTestDAO()
	handler   = franky.NewFrankyHandler(&dao)
	user1     = &franky.User{Id: "123", Email: "email", Password: "password", ApiKey: "apiKey1", CreationDate: "creationDate"}
	user2     = &franky.User{Id: "456", Email: "newEmail", Password: "password", ApiKey: "apiKey2", CreationDate: "creationDate"}
	user3     = &franky.User{Id: "789", Email: "newEmail2", Password: "password", ApiKey: "apiKey3", CreationDate: "creationDate"}
	userJson1 = `{"Id":"123","Email":"email","Password":"password","ApiKey":"apiKey1","CreationDate":"creationDate"}`
	userJson2 = `{"Id":"456","Email":"newEmail","Password":"password","ApiKey":"apiKey2","CreationDate":"creationDate"}`
	userJson3 = `{"Id":"789","Email":"newEmail2","Password":"password","ApiKey":"apiKey3","CreationDate":"creationDate"}`
	records   = `[{"Song":"song","Artist":"artist","Album":"album","Playlist":"playlist","Duration":180,"Time":1000,"Category":"rap"},{"Song":"song","Artist":"artist","Album":"album","Playlist":"playlist","Duration":180,"Time":1000,"Category":"rap"}]`
)

func testHandler(request *http.Request, handler http.HandlerFunc, path string, expectedStatus int, expectedBodyPattern string, test *testing.T) {
	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, handler)
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != expectedStatus {
		test.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	responseBody := strings.TrimSpace(responseRecorder.Body.String())

	matched, err := regexp.MatchString(expectedBodyPattern, responseBody)

	if err != nil {
		test.Errorf("error in testing")
		return
	}

	if !matched {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseBody, expectedBodyPattern)

	}
}
