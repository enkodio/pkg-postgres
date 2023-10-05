package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	cfgEntity "postgres_client/pkg/config/entity"
	"time"
)

const (
	defaultMaxDelay    = 5
	defaultMaxAttempts = 5

	maxConnIdleTime = time.Second * 5
)

type client struct {
	pool        *pgxpool.Pool
	serviceName string
}

func NewClient(cfg cfgEntity.Config, serviceName string) (RepositoryClient, Transactor, error) {
	var pool *pgxpool.Pool
	err := doWithAttempts(
		func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			pgxCfg, err := pgxpool.ParseConfig(cfg.GetDSN(serviceName))
			if err != nil {
				return errors.Wrap(err, "Unable to parse config")
			}
			pgxCfg.MaxConns = int32(cfg.GetMaxOpenConns())
			if pgxCfg.MaxConns == 0 {
				pgxCfg.MaxConns = 4
			}
			pgxCfg.MaxConnIdleTime = maxConnIdleTime
			pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
			if err != nil {
				return errors.Wrap(err, "Failed to connect to postgres")
			}
			return nil
		}, cfg.GetMaxAttempts(), cfg.GetMaxDelay(),
	)

	if err != nil {
		return nil, nil, errors.Wrap(err, "All attempts are exceeded. Unable to connect to postgres")
	}

	pg := client{
		pool:        pool,
		serviceName: serviceName,
	}
	return &pg, &pg, nil
}

func (c *client) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	tx := c.getTx(ctx)
	if tx != nil {
		return tx.Query(ctx, query, args...)
	}
	return c.pool.Query(ctx, query, args...)
}

func (c *client) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	tx := c.getTx(ctx)
	if tx != nil {
		return tx.QueryRow(ctx, query, args...)
	}
	return c.pool.QueryRow(ctx, query, args...)
}

func (c *client) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	tx := c.getTx(ctx)
	if tx != nil {
		return tx.Exec(ctx, query, args...)
	}
	return c.pool.Exec(ctx, query, args...)
}

func doWithAttempts(fn func() error, maxAttempts int, delay int) error {
	var err error
	if maxAttempts == 0 {
		maxAttempts = defaultMaxAttempts
	}
	if delay == 0 {
		delay = defaultMaxDelay
	}
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(time.Second * time.Duration(delay))
			maxAttempts--
			continue
		}
		return nil
	}
	return err
}