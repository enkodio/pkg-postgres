package postgres

import (
	"context"
	"database/sql"
	"io"
)

type RepositoryClient interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) Row
	GetSqlDB() *sql.DB
	CopyFrom(ctx context.Context, file io.ReadCloser, name string) (CommandTag, error)
}

type Transactor interface {
	Begin(*context.Context) error
	Rollback(*context.Context)
	Commit(*context.Context) error
}
