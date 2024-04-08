package client

import (
	cfgEntity "github.com/enkodio/pkg-postgres/internal/pkg/config/entity"
	"github.com/enkodio/pkg-postgres/internal/pkg/logger"
	"github.com/enkodio/pkg-postgres/internal/postgres/logic"
	"github.com/enkodio/pkg-postgres/postgres"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewClient(cfg cfgEntity.Config, serviceName string, log *logrus.Logger) (postgres.RepositoryClient, postgres.Transactor) {
	if log != nil {
		logger.SetLogger(log)
	} else {
		logger.SetDefaultLogger("debug")
	}
	client, err := logic.NewClient(cfg, serviceName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant create pg client"))
	}
	return client, client
}
