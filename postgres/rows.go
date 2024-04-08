package postgres

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
	return pgxscan.ScanAll(dest, r.Rows)
}

// pgx function doesn't work well
// may be try https://github.com/zolstein/pgx-collect
func ScanAll[T any](r Rows) ([]T, error) {
	return pgx.CollectRows(r, pgx.RowToStructByName[T])
}
