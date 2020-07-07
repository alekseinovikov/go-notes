package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitNotesController(router *mux.Router) {
	router.Path("/hello").HandlerFunc(helloHandler)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello! From Notes!"))
}
