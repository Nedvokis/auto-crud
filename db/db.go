package db

import (
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func New() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
