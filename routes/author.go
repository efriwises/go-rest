package routes

import (
	"go-rest/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {

	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")

}
