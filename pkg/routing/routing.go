package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BaseRouter returns all business valuable routes
func BaseRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler()).Methods(http.MethodGet)

	return router
}

func handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requests is processing: %s", r.URL.Path)
		fmt.Fprintf(w, "Congratulations, it works..")
	}
}
