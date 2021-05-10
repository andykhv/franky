package test

import (
	"bytes"
	"fmt"
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
	{User2, UserJson2, http.StatusOK, ""},
	{User1, UserJson1, http.StatusNotFound, "email already exists"},
}

func TestPostUserHandler(tester *testing.T) {
	for _, t := range postUserTests {
		path := fmt.Sprintf("/users/%s", t.user.Id)
		request, err := http.NewRequest("POST", path, bytes.NewBufferString(t.userJson))

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.PostUser, usersRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
