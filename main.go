package main

import (
	"net/http"
	"time"
	"crud_api/api"
	"github.com/gorilla/mux"
	"fmt"
)

var server *http.Server

func init(){
	server = &http.Server{
		Addr: ":8080",
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
}

func main(){
	router := mux.NewRouter()
	api.APIv1(router)

	server.Handler = router
	fmt.Println("Listening on port 8080")
	err := server.ListenAndServe()
	if err != nil{
		panic("server initiation error")
	}
}

