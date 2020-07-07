package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! World!"))
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
