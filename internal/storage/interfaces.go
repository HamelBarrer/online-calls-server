package storage

import "github.com/jackc/pgx/v5"

type SQL interface {
	QueryRow(string, ...interface{}) pgx.Row
	Query(string, ...interface{}) (pgx.Rows, error)
}
