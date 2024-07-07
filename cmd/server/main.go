package main

import (
	"tracking_service/configs"
	"tracking_service/pkg/logger"
	"tracking_service/pkg/server"
)

func main() {
	config := configs.LoadConfig()

	log := logger.NewLogger(config.LogLevel)

	myServer := server.NewServer(config, log)
	myServer.Run()
}
