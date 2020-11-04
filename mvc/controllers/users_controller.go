package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sshindanai/golang-microservices/mvc/services"
	"github.com/sshindanai/golang-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if err != nil {
		utils.RespondError(c, apiErr)
		return
	}

	//return user to client
	utils.Respond(c, http.StatusOK, user)
}
