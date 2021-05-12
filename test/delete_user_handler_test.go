package test

import (
	"fmt"
	"net/http"
	"testing"

	franky "github.com/andykhv/franky/pkg"
)

var deleteUserTests = []struct {
	user           *franky.User
	expectedStatus int
	expectedBody   string
}{
	{user1, http.StatusOK, ""},
	{user2, http.StatusNotFound, fmt.Sprintf("userId %s not found", user2.Id)},
}

func TestDeleteUserHandler(tester *testing.T) {
	for _, t := range deleteUserTests {
		path := fmt.Sprintf("/users/%s", t.user.Id)
		request, err := http.NewRequest("DELETE", path, nil)

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.DeleteUser, usersRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
