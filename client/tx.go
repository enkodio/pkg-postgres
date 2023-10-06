package client

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.enkod.tech/pkg/postgres/pkg/logger"

	"github.com/jackc/pgx/v5"
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
		logger.FromContext(*ctx).Warn(errors.New("empty context"))
		return
	}
	tx := c.getTx(*ctx)
	if tx == nil {
		logger.FromContext(*ctx).Warn(errors.New("empty transaction"))
		return
	}
	err := tx.Rollback(*ctx)
	if err != nil {
		logger.FromContext(*ctx).Warn(err)
	}
	c.deleteTx(ctx)
}

func (c *client) Commit(ctx *context.Context) error {
	if ctx == nil {
		return errors.New("empty context")
	}
	tx := c.getTx(*ctx)
	if tx == nil {
		logger.FromContext(*ctx).Warn(errors.New("empty transaction"))
		return nil
	}
	err := tx.Commit(*ctx)
	if err != nil {
		return errors.Wrap(err, "commit error")
	}
	c.deleteTx(ctx)
	return nil
}
