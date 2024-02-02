package databaselayer

import (
	"errors"
)

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type DinoDBHandler interface {
	GetAvailableDinos() ([]Animal, error)
	GetDinoByNickname(string) (Animal, error)
	GetDinosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"animal_nickname"`
	Zone       int    `bson:"animal_zone"`
	Age        int    `bson:"animal_age"`
}

var DBTypeNotSupported = errors.New("The database type provided is not supported.")

// factory function
func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)
	}
	return nil, DBTypeNotSupported
}
