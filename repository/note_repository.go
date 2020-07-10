package repository

import "github.com/jinzhu/gorm"

type NoteEntity struct {
	ID      int64 `gorm:"primary_key"`
	Title   string
	Content string
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return noteRepository{
		db: db,
	}
}

type NoteRepository interface {
	GetById(id int64) (NoteEntity, bool)
}

func (it noteRepository) GetById(id int64) (NoteEntity, bool) {
	var note NoteEntity

	it.db.First(&note, 1)
	if note.ID == 0 {
		return note, false
	}

	return note, true
}
