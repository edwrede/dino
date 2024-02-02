package main

import (
	"log"

	"github.com/edwrede/dino/databaselayer"
	"github.com/edwrede/dino/dinowebportal/dinoapi"
)

func main() {
	db, err := databaselayer.GetDatabaseHandler(databaselayer.MYSQL, "root:Leighwardo32@@tcp(127.0.0.1:3306)/dino")
	if err != nil {
		log.Fatal(err)
	}
	dinoapi.RunAPI("localhost:8080", db)
}
