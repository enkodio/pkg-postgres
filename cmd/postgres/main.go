package main

import (
	log "github.com/sirupsen/logrus"
	"gitlab.enkod.tech/pkg/postgres/internal/app"
	"gitlab.enkod.tech/pkg/postgres/pkg/config"
	"os"
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
