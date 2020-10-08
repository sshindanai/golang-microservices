package services

import (
	"net/http"

	"github.com/sshindanai/golang-microservices/mvc/domain"
	"github.com/sshindanai/golang-microservices/mvc/utils"
)

type itemsService struct{}

var ItemsService itemsService

func (i *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
	return domain.GetItem(itemId)
}
