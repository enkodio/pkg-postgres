package client

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	cfgEntity "gitlab.enkod.tech/pkg/postgres/pkg/config/entity"
	"gitlab.enkod.tech/pkg/postgres/pkg/logger"
)

func NewClient(cfg cfgEntity.Config, serviceName string, log *logrus.Logger) (RepositoryClient, Transactor) {
	if log != nil {
		logger.SetLogger(log)
	} else {
		logger.SetDefaultLogger("debug")
	}
	client, tx, err := newClient(cfg, serviceName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant create pg client"))
	}
	return client, tx
}
