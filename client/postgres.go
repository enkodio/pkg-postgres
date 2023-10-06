package client

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	postgres "gitlab.enkod.tech/pkg/postgres/internal"
	cfgEntity "gitlab.enkod.tech/pkg/postgres/pkg/config/entity"
)

func NewClient(cfg cfgEntity.Config, serviceName string) (RepositoryClient, Transactor) {
	client, tx, err := postgres.NewClient(cfg, serviceName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant create pg client"))
	}
	return client, tx
}
