package client

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	cfgEntity "gitlab.enkod.tech/pkg/postgres/pkg/config/entity"
)

func NewClient(cfg cfgEntity.Config, serviceName string) (RepositoryClient, Transactor) {
	client, tx, err := newClient(cfg, serviceName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant create pg client"))
	}
	return client, tx
}
