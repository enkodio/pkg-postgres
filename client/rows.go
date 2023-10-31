package client

import "github.com/jackc/pgx/v5"

type Rows struct {
	pgx.Rows
}

func NewRows(rows pgx.Rows, err error) (Rows, error) {
	return Rows{
		Rows: rows,
	}, err
}
