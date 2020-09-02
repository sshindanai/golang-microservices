package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sshindanai/golang-microservices/mvc/services"
	"github.com/sshindanai/golang-microservices/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if err != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
