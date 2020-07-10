package controller

import (
	"github.com/alekseinovikov/go-notes/model"
	"github.com/alekseinovikov/go-notes/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var noteService service.NoteService

func InitNotesController(router *mux.Router, nService service.NoteService) {
	noteService = nService

	router.Path("").Methods("GET").HandlerFunc(listHandler)
	router.Path("").Methods("POST").HandlerFunc(addHandler)
	router.Path("/{id:[0-9]+}").Methods("GET").HandlerFunc(getByIdHandler)
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

func getByIdHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if nil != err {
		log.Fatal(err)
	}

	note, noteError := noteService.FindById(id)
	if nil != noteError {
		HandleError(noteError, writer)
		return
	}

	AddJsonContentHeader(writer.Header())
	WriteAsJson(note, writer)
}
