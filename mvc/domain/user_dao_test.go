package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	if user != nil {
		t.Error("We were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("We were expecting a user with id 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting status 404 not found")
	}
}
