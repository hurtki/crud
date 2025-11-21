package db

import (
	"database/sql"
	"errors"

	pgconn "github.com/jackc/pgx/v5/pgconn"
)

type ErrUniqueViolation struct {
	ConstraintName string
}

func (e ErrUniqueViolation) Error() string {
	return "unique violation with constraint: " + e.ConstraintName
}

type ErrNotNullViolation struct {
	Column string
}

func (e ErrNotNullViolation) Error() string {
	return "not null violation: " + e.Column
}

type ErrSyntaxSql struct {
	Note string
}

func (e ErrSyntaxSql) Error() string {
	return "sql syntax err: " + e.Note
}

var (
	ErrNoRowsAffected      = errors.New("no rows affected")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrNoRows              = errors.New("scan returned error of no rows to scan")
)

func ToDbError(err error) (error, bool) {
	if err == nil {
		return nil, false
	}
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNoRows, true
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return &ErrUniqueViolation{ConstraintName: pgErr.ConstraintName}, true
		case "23503":
			return ErrForeignKeyViolation, true
		case "23502":
			return &ErrNotNullViolation{Column: pgErr.ColumnName}, true
		case "42601":
			return &ErrSyntaxSql{Note: pgErr.Hint}, true
		}
	}
	return err, false
}
