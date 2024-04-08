package postgres

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type Row struct {
	rows   pgx.Rows
	noRows bool
}

func NewRow(rows pgx.Rows) Row {
	return Row{
		rows: rows,
	}
}

func (r *Row) handleErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		r.setNoRows()
		return nil
	}
	return nil
}

func (r *Row) setNoRows() {
	r.noRows = true
}

func (r Row) NoRows() bool {
	return r.noRows
}

func (r Row) Scan(dest ...any) (err error) {
	if err = r.rows.Err(); err != nil {
		return err
	}
	if r.rows.Next() {
		err = r.rows.Scan(dest...)
	} else {
		r.setNoRows()
	}
	r.rows.Close()
	return err
}

func (r Row) ScanStruct(dest interface{}) error {
	if err := r.rows.Err(); err != nil {
		return err
	}
	return r.handleErr(pgxscan.ScanOne(dest, r.rows))
}

// pgx function doesn't work well
// may be try https://github.com/zolstein/pgx-collect
func ScanStruct[T any](r Row) (T, error) {
	value, err := pgx.CollectOneRow(r.rows, pgx.RowToStructByName[T])
	return value, r.handleErr(err)
}
