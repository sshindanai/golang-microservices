package domain

import (
	"fmt"
	"net/http"

	"github.com/sshindanai/golang-microservices/mvc/utils"
)

var (
	item = map[string]*Item{
		"x-100": {ItemId: "x-100", Name: "Stun gun", Price: 100.99},
	}
)

func GetItem(itemId string) (*Item, *utils.ApplicationError) {
	item := item[itemId]

	if item == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Item %v was not found", itemId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return item, nil
}
