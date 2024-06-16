package routes

import (
	"crud_api/api/controllers"
	"github.com/gorilla/mux"
)

func BlogRouter(s *mux.Router) {
	sub_router := s.PathPrefix("/blogs").Subrouter()
	sub_router.HandleFunc("/",controllers.BlogController).Methods("GET")

}