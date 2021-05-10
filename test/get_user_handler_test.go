package test

import (
	"fmt"
	"net/http"
	"testing"
)

var getUserTests = []struct {
	userId         string
	expectedStatus int
	expectedBody   string
}{
	{"123", http.StatusOK, UserJson1},
	{"789", http.StatusNotFound, "userId 789 not found"},
}

func TestGetUserHandler(tester *testing.T) {
	for _, t := range getUserTests {
		path := fmt.Sprintf("/users/%s", t.userId)
		request, err := http.NewRequest("GET", path, nil)

		if err != nil {
			tester.Fatal(err)
		}

		testHandler(request, handler.GetUser, usersRoute, t.expectedStatus, t.expectedBody, tester)
	}
}
