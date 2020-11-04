package domain

import (
	"fmt"
	"net/http"

	"github.com/sshindanai/golang-microservices/mvc/utils"
)

// In memory
var (
	user = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Address: "Bangkok"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := user[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
