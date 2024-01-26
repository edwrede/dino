package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	animal_ID       int
	animal_type     string
	animal_nickname string
	animal_zone     int
	animal_age      int
}

func main() {

	//connect to the database
	//db, err := sql.Open("mysql", "dinouser:LetmeIn123!/dino")
	db, err := sql.Open("mysql", "root:Leighwardo32@@tcp(127.0.0.1:3306)/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Query the db with parameters
	rows, err := db.Query("SELECT * FROM dino.animals WHERE animal_ID > ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}

	for rows.Next() {
		a := animal{}

		err := rows.Scan(&a.animal_ID, &a.animal_type, &a.animal_nickname, &a.animal_zone, &a.animal_age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

	//Insert a row
	//Can use db.Exec for UPDATE or DELETE as well
	//Can also use parameters / query string values like the QUERY above
	result, err := db.Exec("INSERT into dino.animals (animal_type, animal_nickname, animal_zone, animal_age) VALUES ('Carnotausaurus', 'Carno', 3, 22)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
