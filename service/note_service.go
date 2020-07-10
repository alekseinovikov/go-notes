package service

import (
	"github.com/alekseinovikov/go-notes/errors"
	"github.com/alekseinovikov/go-notes/model"
	"github.com/alekseinovikov/go-notes/repository"
)

type noteServiceState struct {
	noteRepository repository.NoteRepository
	lastId         *int64
}

type NoteService interface {
	FindAll() []model.Note
	Add(createRequest model.NoteCreateRequest) model.Note
	Delete(id int64)
	FindById(id int64) (model.Note, errors.NoteError)
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return noteServiceState{
		noteRepository: noteRepository,
	}
}

func (it noteServiceState) FindAll() []model.Note {
	entities := it.noteRepository.FindAll()
	return convertEntities(entities)
}

func (it noteServiceState) FindById(id int64) (model.Note, errors.NoteError) {
	entity, found := it.noteRepository.FindById(id)
	if !found {
		return model.Note{}, errors.NotFoundError
	}

	return convertEntityToModel(entity), nil
}

func (it noteServiceState) Add(request model.NoteCreateRequest) model.Note {
	entity := convertAddRequestToEntity(request)
	savedEntity := it.noteRepository.Save(entity)

	return convertEntityToModel(savedEntity)
}

func (it noteServiceState) Delete(id int64) {
	it.noteRepository.Delete(id)
}

func convertEntities(entities []repository.NoteEntity) []model.Note {
	result := make([]model.Note, 0)

	for _, entity := range entities {
		note := convertEntityToModel(entity)
		result = append(result, note)
	}

	return result
}

func convertEntityToModel(noteEntity repository.NoteEntity) model.Note {
	return model.Note{
		Id:      noteEntity.ID,
		Title:   noteEntity.Title,
		Content: noteEntity.Content,
	}
}

func convertAddRequestToEntity(request model.NoteCreateRequest) repository.NoteEntity {
	return repository.NoteEntity{
		Title:   request.Title,
		Content: request.Content,
	}
}
