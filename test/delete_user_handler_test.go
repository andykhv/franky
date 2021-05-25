package test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

var deleteUserTests = []struct {
	user           *franky.User
	apiKey         string
	expectedStatus int
	expectedBody   string
}{
	{user1, user1.ApiKey, http.StatusOK, ""},
	{user2, user2.ApiKey, http.StatusNotFound, fmt.Sprintf("userId %s not found", user2.Id)},
	{user1, "wrongApiKey", http.StatusUnauthorized, "invalid token"},
}

func TestDeleteUserHandler(tester *testing.T) {
	for _, t := range deleteUserTests {
		path := fmt.Sprintf("/users/%s", t.user.Id)
		request, err := http.NewRequest("DELETE", path, bytes.NewBufferString(fmt.Sprintf("{\"token\":\"%s\"}", t.apiKey)))

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.DeleteUser, usersIdRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
