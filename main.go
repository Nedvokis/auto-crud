package main

import (
	"fmt"
	"os"

	"github.com/auto-crud/db"
	"github.com/auto-crud/models"
)

func main() {
	some := &models.User{}
	some.Name = "Denis"
	some.SecondName = "Trupper"
	some.Online = true
	some.CreatedAt = "2022-03-04"
	some.UpdatedAt = "2022-03-04"

	newDb, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	store := db.NewStore(newDb)
	store.Create("user", some)
}
