package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

var getUserTests = []struct {
	user           *franky.User
	apiKey         string
	expectedStatus int
	expectedBody   string
}{
	{user1, user1.ApiKey, http.StatusOK, userJson1},
	{user1, "wrong api key", http.StatusUnauthorized, "invalid token"},
	{user3, user3.ApiKey, http.StatusNotFound, "userId 789 not found"},
}

func TestGetUserHandler(tester *testing.T) {
	for _, t := range getUserTests {
		path := fmt.Sprintf("/users/%s", t.user.Id)
		request, err := http.NewRequest("GET", path, bytes.NewBufferString(fmt.Sprintf("{\"token\":\"%s\"}", t.apiKey)))

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.GetUser, usersIdRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
