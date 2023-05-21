package main

import (
	"github.com/Luan-max/go-jobs/config"
	"github.com/Luan-max/go-jobs/router"
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

	router.Initialize()
}
