package services

import (
	"github.com/sshindanai/golang-microservices/mvc/domain"
	"github.com/sshindanai/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
