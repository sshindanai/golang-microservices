package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sshindanai/golang-microservices/src/api/api/errors"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
	"github.com/sshindanai/golang-microservices/src/api/services"
)

func CreateRepo(c *gin.Context) {
	var req repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(req)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var req []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(req)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}
