package controller

import (
	"github.com/alekseinovikov/go-notes/model"
	"github.com/alekseinovikov/go-notes/service"
	"github.com/gorilla/mux"
	"net/http"
)

var noteService = service.NewNoteService()

func InitNotesController(router *mux.Router) {
	router.Path("").Methods("GET").HandlerFunc(listHandler)
	router.Path("").Methods("POST").HandlerFunc(addHandler)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	AddJsonContentHeader(writer.Header())

	notes := noteService.FindAll()
	WriteAsJson(notes, writer)
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	var createRequest model.NoteCreateRequest
	FromJsonRequest(request, &createRequest)

	newNote := noteService.Add(createRequest)

	AddJsonContentHeader(writer.Header())
	WriteAsJson(newNote, writer)
}
