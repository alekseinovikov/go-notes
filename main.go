package main

import (
	"github.com/alekseinovikov/go-notes/controller"
	"github.com/alekseinovikov/go-notes/repository"
	"github.com/alekseinovikov/go-notes/service"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

func main() {
	db := initDB()
	defer db.Close()

	noteRepository := initNoteRepository(db)
	noteService := initNoteService(noteRepository)

	// Create Server and Route Handlers
	r := mux.NewRouter().StrictSlash(false)

	notesSubRouter := r.PathPrefix("/notes").Subrouter()
	controller.InitNotesController(notesSubRouter, noteService)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		log.Fatal(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "note.db")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func initNoteRepository(db *gorm.DB) repository.NoteRepository {
	return repository.NewNoteRepository(db)
}

func initNoteService(noteRepository repository.NoteRepository) service.NoteService {
	return service.NewNoteService(noteRepository)
}
