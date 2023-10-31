package app

import (
	"context"
	"github.com/enkodio/pkg-postgres/client"
	cfgEntity "github.com/enkodio/pkg-postgres/pkg/config/entity"
	"github.com/enkodio/pkg-postgres/pkg/logger"
)

func Run(configSettings cfgEntity.Settings, serviceName string) {
	var (
		pgClient, transactor = client.NewClient(configSettings.PostgresConfigs, serviceName, nil)
	)
	log := logger.GetLogger()
	ctx := context.Background()
	err := transactor.Begin(&ctx)
	if err != nil {
		log.WithError(err).Error("cant begin transaction")
		return
	}

	_, err = pgClient.Exec(ctx, `CREATE TABLE IF NOT EXISTS test(
    id int8 NOT NULL,
    PRIMARY KEY (id)
);`)
	if err != nil {
		log.WithError(err).Error("cant create table")
		return
	}

	_, err = pgClient.Exec(ctx, "INSERT INTO test VALUES ($1)", 14)
	if err != nil {
		log.WithError(err).Error("cant insert value")
	}

	err = transactor.Commit(&ctx)
	if err != nil {
		log.WithError(err).Error("cant commit transaction")
	}

	var id int64
	err = pgClient.QueryRow(ctx, "SELECT id FROM test WHERE id = 13").Scan(&id)
	if err != nil {
		log.WithError(err).Error("cant exec sql request")
		return
	}

	rows, err := pgClient.Query(ctx, "SELECT id FROM test")
	if err != nil {
		log.WithError(err).Error("cant exec sql request")
		return
	}
	defer rows.Close()
	var ids []int64
	err = rows.ScanAll(&ids)
	if err != nil {
		log.WithError(err).Error("cant scan values")
		return
	}
}
