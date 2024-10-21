package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/config"
	"gitlab.com/dert-ops/mediCat/mediCat-Dev.git/cmd/routes"
)

func main() {
	config.LoadEnv()

	config.InitLogrusLogger()

	router := gin.Default()
	routes.UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.LogrusLogger.Infof("User Service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		config.LogrusLogger.Fatalf("Failed to run server: %v", err)
	}
}
