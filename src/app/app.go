package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sshindanai/golang-microservices/src/log/optionb"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	optionb.Info("about to map the urls", optionb.Field("step", "1"), optionb.Field("status", "pending"))
	mapUrls()
	optionb.Info("url successfully", optionb.Field("step", "2"), optionb.Field("status", "success"))

	if err := router.Run(":8080"); err != nil {
		optionb.Error("router is down", err, optionb.Field("status", "failed"))
	}
}
