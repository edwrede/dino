package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) GetAvailableDinos() ([]Animal, error) {
	return handler.sendQuery("SELECT * FROM dino.animals")
}

func (handler *SQLHandler) GetDinoByNickname(nickname string) (Animal, error) {
	row := handler.QueryRow(fmt.Sprintf("SELECT * FROM dino.animals WHERE animal_nickname = '%s'", nickname)) //This gives support for all SQL databases and is db engine agnostic

	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}

func (handler *SQLHandler) GetDinosByType(dinoType string) ([]Animal, error) {

	//Let's check the exact query and log it out to make sure it's as we expect
	log.Println(fmt.Sprintf("SELECT * FROM dino.animals WHERE animal_type = '%s'", dinoType))
	return handler.sendQuery(fmt.Sprintf("SELECT * FROM dino.animals WHERE animal_type = '%s'", dinoType))
}

func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("INSERT INTO dino.animals (animal_type, animal_nickname, animal_zone, animal_age) VALUES('%s', '%s', %d, %d)", a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}

func (handler *SQLHandler) UpdateAnimal(a Animal, animalNickname string) error {
	_, err := handler.Exec(fmt.Sprintf("UPDATE dino.animals SET animal_type = '%s', animal_nickname = '%s', animal_zone = %d, animal_age = %d WHERE animal_nickname = '%s'", a.AnimalType, a.Nickname, a.Zone, a.Age, animalNickname))
	return err
}

// This function allows us to return query rows as animal types
func (handler *SQLHandler) sendQuery(query string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}
