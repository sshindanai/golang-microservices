package app

import (
	"github.com/sshindanai/golang-microservices/src/api/api/controllers/polo"
	"github.com/sshindanai/golang-microservices/src/api/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
