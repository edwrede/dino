package dinoapi

import (
	"net/http"

	"github.com/edwrede/dino/databaselayer"
	"github.com/gorilla/mux"
)

func RunAPI(endpoint string, db databaselayer.DinoDBHandler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db databaselayer.DinoDBHandler) {
	handler := newDinoRESTAPIHandler(db)
	apiRouter := r.PathPrefix("/api/dinos").Subrouter()

	apiRouter.Methods("GET").Path("/{searchCriteria}/{search}").HandlerFunc((handler.searchHandler))
	apiRouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)

}
