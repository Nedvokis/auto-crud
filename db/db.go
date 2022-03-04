package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
