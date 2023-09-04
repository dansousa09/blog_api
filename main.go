package main

import (
	"github.com/dansousa09/blog_api/config"
	"github.com/dansousa09/blog_api/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
	}

	router.Initialize()
}
