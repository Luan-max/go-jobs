package main

import (
	"github.com/Luan-max/go-jobs/config"
	"github.com/Luan-max/go-jobs/router"
	"github.com/joho/godotenv"
)

var (
	logger config.Logger
)

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
