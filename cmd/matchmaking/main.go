package main

import (
	"log"
	"matchmaking/config"
	"matchmaking/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	// Configuration
	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
