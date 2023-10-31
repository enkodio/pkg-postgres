package client

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type Rows struct {
	pgx.Rows
}

func NewRows(rows pgx.Rows, err error) (Rows, error) {
	return Rows{
		Rows: rows,
	}, err
}

func (r Rows) ScanAll(dest interface{}) error {
	err := pgxscan.ScanAll(dest, r.Rows)
	if err != nil {
		return err
	}
	return nil
}
