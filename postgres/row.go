package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

type Row struct {
	pgx.Row
	noRows bool
}

func NewRow(row pgx.Row) Row {
	return Row{
		Row: row,
	}
}

func (r *Row) setNoRows() {
	r.noRows = true
}

func (r Row) NoRows() bool {
	return r.noRows
}

func (r Row) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if errors.Is(err, pgx.ErrNoRows) {
		r.setNoRows()
		return nil
	}
	return err
}
