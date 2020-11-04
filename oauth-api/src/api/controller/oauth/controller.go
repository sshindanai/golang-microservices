package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sshindanai/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/sshindanai/golang-microservices/oauth-api/src/api/services"
	"github.com/sshindanai/golang-microservices/src/api/api/errors"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
