package postgres

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	postgres "postgres_client/internal"
	cfgEntity "postgres_client/pkg/config/entity"
)

func NewClient(cfg cfgEntity.Config, serviceName string) (postgres.RepositoryClient, postgres.Transactor) {
	client, tx, err := postgres.NewClient(cfg, serviceName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant create pg client"))
	}
	return client, tx
}
