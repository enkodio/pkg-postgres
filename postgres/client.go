package postgres

import (
	"context"
	"database/sql"
)

type RepositoryClient interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	GetSqlDB() *sql.DB
}

type Transactor interface {
	Begin(*context.Context) error
	Rollback(ctx *context.Context)
	Commit(ctx *context.Context) error
}
