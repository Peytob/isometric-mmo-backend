package main

import (
	"isonetric-mmo-backend/configs"
	log "log/slog"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	if err := configs.InitLogging(config); err != nil {
		panic(err)
	}

	log.Info("Configuration loaded. Initializing application")
}
