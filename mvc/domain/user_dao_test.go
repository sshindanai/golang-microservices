package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialize
	userID := 0
	// Execution
	user, err := GetUser(int64(userID))

	// Validation
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting a user with id 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUseNoError(t *testing.T) {
	userID := 123

	user, err := GetUser(int64(userID))

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
}
