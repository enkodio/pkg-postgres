package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"postgres_client/internal/app"
	"postgres_client/pkg/config"
)

const (
	serviceName = "test_kafka_client"
)

func main() {
	configSettings, err := config.LoadConfigSettingsByPath("configs")
	if err != nil {
		log.Error(err)
		os.Exit(2)
	}
	app.Run(configSettings, serviceName)
}
