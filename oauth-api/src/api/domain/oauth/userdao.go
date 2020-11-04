package oauth

import (
	"fmt"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?"
)

// Mock db
var (
	users = map[string]*User{
		"shin": {Id: 123, Username: "shin"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user with nickname '%s'", username))
	}

	return users[username], nil
}
