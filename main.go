package main

import (
	"github.com/Luan-max/go-jobs/config"
	"github.com/Luan-max/go-jobs/router"
	"github.com/joho/godotenv"
)

var (
	logger config.Logger
)

// @title           Payment Gateway
// @version         1.0
// @description     This is a sample payment gateway.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("Initialize error: %v", err)
		return
	}

	err = godotenv.Load()
	if err != nil {
		logger.Errf("Error loading .env file: %v", err)
		return
	}

	router.Initialize()
}
