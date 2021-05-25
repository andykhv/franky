package test

import (
	"bytes"
	"net/http"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

var postUserTests = []struct {
	user           *franky.User
	userJson       string
	expectedStatus int
	expectedBody   string
}{
	{user2, userJson2, http.StatusOK, `{\\\"token\\\":\\\"franky.api.[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\\\", \\\"id\\\":\\\"franky.user.[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\\\"}`},
	{user1, userJson1, http.StatusNotFound, "email already exists"},
}

func TestPostUserHandler(tester *testing.T) {
	for _, t := range postUserTests {
		path := "/users"
		request, err := http.NewRequest("POST", path, bytes.NewBufferString(t.userJson))

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.PostUser, usersRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
