package app

import (
	"github.com/davidlamyc/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Starting the application...")
	router.Run(":8080")
}