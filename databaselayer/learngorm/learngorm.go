package main

import (
	"log"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type animal struct {
	//gorm model
	animal_ID       int    `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	animal_type     string `gorm:"type:TEXT"`
	animal_nickname string `gorm:"type:TEXT"`
	animal_zone     int    `gorm:"type:INTEGER"`
	animal_age      int    `gorm:"type:INTEGER"`
}

func getAnimalByID(db *gorm.DB, animalID uint) (*animal, error) {
	var myAnimal animal
	result := db.First(&myAnimal, animalID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &myAnimal, nil
}

func main() {

	dsn := "root:Leighwardo32@@tcp(127.0.0.1:3306)/dino?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db.Attrs())

	// Perform database migration
	err = db.AutoMigrate(&animal{})
	if err != nil {
		log.Fatal(err)
	}

	// Query user by ID
	animal, err := getAnimalByID(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Animal by ID:", animal)

}
