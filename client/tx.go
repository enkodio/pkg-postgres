package client

import (
	"context"
	"github.com/enkodio/pkg-postgres/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	txKey = "txKey"
)

func (c *client) deleteTx(ctx *context.Context) {
	*ctx = context.WithValue(*ctx, txKey+c.serviceName, nil)
}

func (c *client) setTx(ctx *context.Context, tx pgx.Tx) {
	*ctx = context.WithValue(*ctx, txKey+c.serviceName, tx)
}

func (c *client) getTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey + c.serviceName).(pgx.Tx); ok {
		return tx
	}
	return nil
}

func (c *client) Begin(ctx *context.Context) error {
	if ctx == nil {
		return errors.New("empty context")
	}
	tx, err := c.pool.Begin(*ctx)
	if err != nil {
		return err
	}
	c.setTx(ctx, tx)
	return nil
}

func (c *client) Rollback(ctx *context.Context) {
	if ctx == nil {
		logger.GetLogger().Warn(errors.New("empty context"))
		return
	}
	tx := c.getTx(*ctx)
	if tx == nil {
		logger.GetLogger().Warn(errors.New("empty transaction"))
		return
	}
	err := tx.Rollback(*ctx)
	if err != nil {
		logger.GetLogger().Warn(err)
	}
	c.deleteTx(ctx)
}

func (c *client) Commit(ctx *context.Context) error {
	if ctx == nil {
		return errors.New("empty context")
	}
	tx := c.getTx(*ctx)
	if tx == nil {
		logger.GetLogger().Warn(errors.New("empty transaction"))
		return nil
	}
	defer c.deleteTx(ctx)
	err := tx.Commit(*ctx)
	if err != nil {
		return errors.Wrap(err, "commit error")
	}
	return nil
}
