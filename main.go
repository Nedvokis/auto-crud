package main

import (
	"fmt"
	"os"

	"github.com/auto-crud/db"
	"github.com/auto-crud/models"
	gofakeit "github.com/brianvoe/gofakeit/v6"
)

func main() {
	some := &models.InsertUser{}
	some.Name = gofakeit.Name()
	some.SecondName = "Trupper"
	some.Online = false
	some.CreatedAt = "2022-03-04"
	some.UpdatedAt = "2022-03-04"

	newDb, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	store := db.NewStore(newDb)
	if err := store.Create("user", some); err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	result, err := store.GetOne("user", 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(result)

}
