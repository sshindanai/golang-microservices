package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapURLs()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
