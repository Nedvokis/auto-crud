package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/auto-crud/models"
	"github.com/fatih/structs"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (scs *Store) Create(collectionName string, collection interface{}) error {
	ctx := context.Background()

	coll := structs.Map(collection)

	var keys, values []string

	for k, v := range coll {
		keys = append(keys, k)

		var valueString string

		switch v.(type) {
		case string:
			valueString = fmt.Sprintf("'%v'", v)
		case int:
			valueString = fmt.Sprintf("%v", v)
		default:
			valueString = fmt.Sprintf("%v", v)
		}

		values = append(values, valueString)
	}

	keysStr := strings.Join(keys, ", ")
	valuesStr := strings.Join(values, ", ")

	query := fmt.Sprintf(INSERT, collectionName, keysStr, valuesStr)
	scs.db.QueryRow(ctx, query)

	return nil
}

func (scs *Store) GetOne(collectionName string, id uint) (interface{}, error) {
	ctx := context.Background()

	user := models.SelectUser{}

	query := fmt.Sprintf(READ_ONE, collectionName, id)
	if err := scs.db.QueryRow(ctx, query).Scan(
		&user.Id,
		&user.Name,
		&user.SecondName,
		&user.Online,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (scs *Store) GetAll(collectionName string, collection interface{}) error {
	ctx := context.Background()
	coll := structs.Map(collection)

	var where []string

	for k, v := range coll {
		fmt.Println(k, "some", v)

		var valueString string

		switch v.(type) {
		case string:
			valueString = fmt.Sprintf("'%v'", v)
		case int:
			valueString = fmt.Sprintf("%v", v)
		default:
			valueString = fmt.Sprintf("%v", v)
		}

		where = append(where, fmt.Sprintf("%s=%s", k, valueString))
	}

	whereStr := strings.Join(where, " and ")

	query := fmt.Sprintf(READ_ALL, collectionName, whereStr)
	if err := scs.db.QueryRow(ctx, query).Scan(&collection); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	return nil
}

func (scs *Store) Update() error {
	return nil
}

func (scs *Store) Delete() error {
	return nil
}
