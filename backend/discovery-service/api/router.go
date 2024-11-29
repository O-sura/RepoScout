package api

import (
	"example.com/discovery-service/internal"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/repositories", internal.GetFilteredRepositories).Methods("GET")
	router.HandleFunc("/repository/{id}", internal.GetSingleRepository).Methods("GET")
	return router
}
