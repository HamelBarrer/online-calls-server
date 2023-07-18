package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type postgres struct {
	p *pgxpool.Pool
}

func NewPostgres() (SQL, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	databaseUser := viper.Get("DATABASE_USER")
	databasePassword := viper.Get("DATABASE_PASSWORD")
	databaseHost := viper.Get("DATABASE_HOST")
	databasePort := viper.Get("DATABASE_PORT")
	databaseDB := viper.Get("DATABASE_DB")

	connUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseDB)

	p, err := pgxpool.New(context.Background(), connUrl)
	if err != nil {
		return nil, err
	}

	return &postgres{p}, nil
}

func (p *postgres) QueryRow(query string, args ...interface{}) pgx.Row {
	return p.p.QueryRow(context.Background(), query, args...)
}

func (p *postgres) Query(query string, args ...interface{}) (pgx.Rows, error) {
	return p.p.Query(context.Background(), query, args...)
}
