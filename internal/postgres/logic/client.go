package logic

import (
	"context"
	"database/sql"
	cfgEntity "github.com/enkodio/pkg-postgres/internal/pkg/config/entity"
	"github.com/enkodio/pkg-postgres/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"time"
)

type Client struct {
	pool        *pgxpool.Pool
	pgxCfg      *pgxpool.Config
	serviceName string
}

func NewClient(cfg cfgEntity.Config, serviceName string) (*Client, error) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.GetDSN(serviceName))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse config")
	}
	pgxCfg.MaxConns = int32(cfg.GetMaxOpenConns())
	if pgxCfg.MaxConns == 0 {
		pgxCfg.MaxConns = defaultMaxOpenConns
	}
	pgxCfg.MaxConnIdleTime = maxConnIdleTime

	var pool *pgxpool.Pool

	pg := Client{
		serviceName: serviceName,
		pgxCfg:      pgxCfg,
	}

	err = doWithAttempts(
		func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
			if err != nil {
				return errors.Wrap(err, "Failed to connect to postgres")
			}
			return nil
		}, cfg.GetMaxAttempts(), cfg.GetMaxDelay(),
	)

	if err != nil {
		return nil, errors.Wrap(err, "All attempts are exceeded. Unable to connect to postgres")
	}

	pg.pool = pool
	return &pg, nil
}

func (c *Client) Query(ctx context.Context, query string, args ...interface{}) (postgres.Rows, error) {
	tx := c.getTx(ctx)
	if tx != nil {
		return postgres.NewRows(tx.Query(ctx, query, args...))
	}
	return postgres.NewRows(c.pool.Query(ctx, query, args...))
}

func (c *Client) QueryRow(ctx context.Context, query string, args ...interface{}) postgres.Row {
	tx := c.getTx(ctx)
	if tx != nil {
		return postgres.NewRow(tx.QueryRow(ctx, query, args...))
	}
	return postgres.NewRow(c.pool.QueryRow(ctx, query, args...))
}

func (c *Client) Exec(ctx context.Context, query string, args ...interface{}) (postgres.CommandTag, error) {
	tx := c.getTx(ctx)
	if tx != nil {
		return postgres.NewCommandTag(tx.Exec(ctx, query, args...))
	}
	return postgres.NewCommandTag(c.pool.Exec(ctx, query, args...))
}

func (c *Client) GetSqlDB() *sql.DB {
	return stdlib.OpenDBFromPool(c.pool)
}
