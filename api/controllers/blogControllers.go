package controllers

import (
	"net/http"
	"io"
)

func BlogController(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w,"Hello world\n")
}
	