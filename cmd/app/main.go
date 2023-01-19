package main

import (
	"log"

	"github.com/miromax42/dots-backend/config"
	"github.com/miromax42/dots-backend/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
