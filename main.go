package main

import (
	"github.com/alekseinovikov/go-notes/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter().StrictSlash(false)

	notesSubRouter := r.PathPrefix("/notes").Subrouter()
	controller.InitNotesController(notesSubRouter)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		log.Fatal(err)
	}
}
