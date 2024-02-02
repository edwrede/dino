package dinoapi

import (
	"encoding/json"
	"fmt"
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

	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found`)
		return
	}

	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error

	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDinoByNickname(searchKey)
	case "type":
		animals, err = handler.dbhandler.GetDinosByType(searchKey)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occurred while querying animals %v ", err)
	}
	if len(animals) > 0 {
		json.NewEncoder(w).Encode(animals)
		return
	}
	json.NewEncoder(w).Encode(animal)
}

func (handler *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found`)
		return
	}

	searchKey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `No search criteria found`)
		return
	}

	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error

	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetDinoByNickname(searchKey)
	case "type":
		animals, err = handler.dbhandler.GetDinosByType(searchKey)
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occurred while querying animals %v ", err)
	}
	if len(animals) > 0 {
		json.NewEncoder(w).Encode(animals)
		return
	}
	json.NewEncoder(w).Encode(animal)
}
