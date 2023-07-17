package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Postgres struct {
	p *pgxpool.Pool
}

func NewPostgres() (*Postgres, error) {
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

	return &Postgres{p}, nil
}
