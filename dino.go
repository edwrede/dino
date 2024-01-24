package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/edwrede/dino/dinowebportal"
)

type configuration struct {
	WebServer string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	dinowebportal.RubWebPortal(config.WebServer)
}
