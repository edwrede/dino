package dinoapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/edwrede/dino/databaselayer"
	"github.com/gorilla/mux"
)

type DinoRESTAPIHandler struct {
	dbhandler databaselayer.DinoDBHandler
}

func newDinoRESTAPIHandler(db databaselayer.DinoDBHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {

	//Here we are validating our code to make sure that there is a value in the map allocated to SearchCriteria, if not the search criteria path will not be triggered
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found`)
		return
	}

	//Similar to above if we've found SearchCriteria, but we don't find any value to search for, we don't make the request
	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found`)
		return
	}

	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error

	//Now we evaluate the type of search, whether it is by nickname or by type. Nickname returns a single row, type returns multiple rows
	//The searchKey value was set previously if the "search" key is found in the map, then searchKey contains the value
	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDinoByNickname(searchKey)
	case "type":
		animals, err = handler.dbhandler.GetDinosByType(searchKey)
		if len(animals) > 0 {
			//HTTP response writer is passed into a JSON encoder and the animals array is encoded into the response writer
			json.NewEncoder(w).Encode(animals)
			return
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occurred while querying animals %v ", err)
	}

	//If we get here animals was length 0 and this is a single animal return
	json.NewEncoder(w).Encode(animal)
}

func (handler *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `Operation was not provided, please use add or edit`)
		return
	}

	var animal databaselayer.Animal

	//The POST object comes with a body that contains the data passed with the POST request
	//REFERENCE to the animal variable just created is passed into the decoder as the decoder only accepts a reference
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "could not decode the request body to json %v", err)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		//This extracts nickname from the URI as the URI will look like /api/dinos/edit/rex
		nickname := r.RequestURI[len("/api/dinos/edit/"):]
		log.Println("edit request for nickname: ", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occurred while processing request %v", err)
	}
}
