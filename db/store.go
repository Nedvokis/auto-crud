package db

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/structs"
	pgx "github.com/jackc/pgx/v4"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{
		db: db,
	}
}

func (scs *Store) Create(collectionName string, collection struct{}) error {
	ctx := context.Background()
	coll := structs.Map(collection)
	fmt.Println(coll)
	for k, v := range coll {
		fmt.Println(k, v)
	}
	query := fmt.Sprintf(INSERT, collectionName)
	var greeting string
	if err := scs.db.QueryRow(ctx, query).Scan(&greeting); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}

	return nil
}

func (scs *Store) GetByID() error {
	return nil
}

func (scs *Store) GetByUserID() error {
	return nil
}

func (scs *Store) Update() error {
	return nil
}

func (scs *Store) Delete() error {
	return nil
}
