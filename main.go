package main

import (
	"fmt"
	"os"

	"github.com/auto-crud/db"
	"github.com/auto-crud/handler"
	"github.com/auto-crud/router"
)

func main() {
	r := router.New()

	v1 := r.Group("/api/auto-crud")

	newDb, err := db.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	store := db.NewStore(newDb)
	h := handler.NewHandler(store)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8100"))

	// var user models.SelectUser
	// result, err := store.GetOne("user", 1, user)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(result)

}
