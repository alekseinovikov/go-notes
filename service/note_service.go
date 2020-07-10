package service

import (
	"github.com/alekseinovikov/go-notes/errors"
	"github.com/alekseinovikov/go-notes/model"
	"github.com/alekseinovikov/go-notes/repository"
	"sync/atomic"
)

type noteServiceState struct {
	noteRepository repository.NoteRepository
	notes          map[int64]model.Note
	lastId         *int64
}

type NoteService interface {
	FindAll() []model.Note
	Add(createRequest model.NoteCreateRequest) model.Note
	Delete(id int64)
	FindById(id int64) (model.Note, errors.NoteError)
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	var initId int64 = 0
	return noteServiceState{
		noteRepository: noteRepository,
		notes:          make(map[int64]model.Note),
		lastId:         &initId,
	}
}

func (it noteServiceState) FindAll() []model.Note {
	values := make([]model.Note, 0, len(it.notes))
	for _, v := range it.notes {
		values = append(values, v)
	}

	return values
}

func (it noteServiceState) FindById(id int64) (model.Note, errors.NoteError) {
	note, found := it.notes[id]
	if !found {
		return model.Note{}, errors.NotFoundError
	}

	return note, nil
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
