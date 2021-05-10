package test

import (
	"net/http"
	"net/http/httptest"
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
	User1     = &franky.User{Id: "123", Email: "email", Password: "password", ApiKey: "apiKey", CreationDate: "creationDate"}
	User2     = &franky.User{Id: "456", Email: "newEmail", Password: "password", ApiKey: "apiKey", CreationDate: "creationDate"}
	UserJson1 = `{"Id":"123","Email":"email","Password":"password","ApiKey":"apiKey","CreationDate":"creationDate"}`
	UserJson2 = `{"Id":"456","Email":"newEmail","Password":"password","ApiKey":"apiKey","CreationDate":"creationDate"}`
	Records   = `[{"Song":"song","Artist":"artist","Album":"album","Playlist":"playlist","Duration":180,"Time":1000,"Category":"rap"},{"Song":"song","Artist":"artist","Album":"album","Playlist":"playlist","Duration":180,"Time":1000,"Category":"rap"}]`
)

func testHandler(request *http.Request, handler http.HandlerFunc, path string, expectedStatus int, expectedBody string, test *testing.T) {
	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, handler)
	router.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != expectedStatus {
		test.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	if responseBody := strings.TrimSpace(responseRecorder.Body.String()); responseBody != expectedBody {
		test.Errorf("handler returned unexpected body: got %s want %s",
			responseBody, expectedBody)
	}
}
