package dinowebportal

import (
	"fmt"
	"net/http"
)

func RubWebPortal(webServerAddress string) error {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(webServerAddress, nil)
	return err
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Dino web portal %s", request.RemoteAddr)
}
