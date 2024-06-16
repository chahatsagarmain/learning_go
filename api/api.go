package api

import (
	"github.com/gorilla/mux"
	"crud_api/api/routes"
)

func APIv1(r *mux.Router) *mux.Router{
	s := r.PathPrefix("/api/v1").Subrouter()
	routes.BlogRouter(s)
	return s
}