package logic

import (
	"context"
	"github.com/enkodio/pkg-postgres/internal/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	txKey = "txKey"
)

func (c *Client) deleteTx(ctx *context.Context) {
	*ctx = context.WithValue(*ctx, txKey+c.serviceName, nil)
}

func (c *Client) setTx(ctx *context.Context, tx pgx.Tx) {
	*ctx = context.WithValue(*ctx, txKey+c.serviceName, tx)
}

func (c *Client) getTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey + c.serviceName).(pgx.Tx); ok {
		return tx
	}
	return nil
}

func (c *Client) Begin(ctx *context.Context) error {
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

func (c *Client) Rollback(ctx *context.Context) {
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

func (c *Client) Commit(ctx *context.Context) error {
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
