package main

import (
	"fmt"
	"time"

	"github.com/auto-crud/models"
	"github.com/fatih/structs"
)

func main() {
	some := &models.User{}
	some.Id = 1
	some.Name = "Denis"
	some.SecondName = "Trupper"
	some.Online = true
	some.CreatedAt = time.Now()
	some.UpdatedAt = time.Now()
	coll := structs.Map(some)
	fmt.Println(coll)
	for k, v := range coll {
		fmt.Println(k, "some", v)
	}
}
