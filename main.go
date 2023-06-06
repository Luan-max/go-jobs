package main

import (
	"github.com/Luan-max/go-jobs/infra/config"
	"github.com/Luan-max/go-jobs/infra/database"
	"github.com/Luan-max/go-jobs/infra/http/router"
	"github.com/joho/godotenv"
)

func main() {
	logger := config.GetLogger("main")

	err := database.Init()
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
