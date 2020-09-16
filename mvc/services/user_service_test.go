package services

import (
	"net/http"
	"testing"

	"github.com/sshindanai/golang-microservices/mvc/domain"

	"github.com/sshindanai/golang-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock
}

// Mock region
type usersDaoMock struct{}

// Mock GetUser function
func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 doesn't exists",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 doesn't exists", err.Message)
}

func TestGetUserNotFoundNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123, FirstName: "Fede", LastName: "Leon", Address: "Bangkok"}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
