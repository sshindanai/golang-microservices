package app

import (
	"github.com/sshindanai/golang-microservices/oauth-api/src/api/controller/oauth"
	"github.com/sshindanai/golang-microservices/src/api/api/controllers/polo"
)

func mapURLs() {
	// Register api router
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}

//curl -X POST "localhost:8080/oauth/accesstoken" -d '{"username": "shin", "password": "testing"}'
