package main

import (
	"github.com/enkodio/pkg-postgres/internal/pkg/config"
	"github.com/enkodio/pkg-postgres/internal/postgres/app"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	serviceName = "test_postgres_client"
)

func main() {
	configSettings, err := config.LoadConfigSettingsByPath("internal/cmd/configs")
	if err != nil {
		log.Error(err)
		os.Exit(2)
	}
	app.Run(configSettings, serviceName)
}
