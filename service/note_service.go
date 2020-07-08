package service

import (
	"github.com/alekseinovikov/go-notes/model"
	"sync/atomic"
)

type noteServiceState struct {
	notes  map[int64]model.Note
	lastId *int64
}

type NoteService interface {
	FindAll() []model.Note
	Add(createRequest model.NoteCreateRequest) model.Note
	Delete(id int64)
	Get(id int64) model.Note
}

func NewNoteService() NoteService {
	var initId int64 = 0
	return noteServiceState{
		notes:  make(map[int64]model.Note),
		lastId: &initId,
	}
}

func (it noteServiceState) FindAll() []model.Note {
	values := make([]model.Note, 0, len(it.notes))
	for _, v := range it.notes {
		values = append(values, v)
	}

	return values
}

func (it noteServiceState) Get(id int64) model.Note {
	return it.notes[id]
}

func (it noteServiceState) Add(request model.NoteCreateRequest) model.Note {
	newId := it.getNextId()

	newNote := model.Note{
		Id:      newId,
		Title:   request.Title,
		Content: request.Content,
	}

	it.notes[newId] = newNote
	return newNote
}

func (it noteServiceState) Delete(id int64) {
	delete(it.notes, id)
}

func (it noteServiceState) getNextId() int64 {
	return atomic.AddInt64(it.lastId, 1)
}

func (it noteServiceState) addNote(note model.Note) {
	it.notes[note.Id] = note
}
