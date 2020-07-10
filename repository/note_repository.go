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
	migrateSchema(db)

	return noteRepository{
		db: db,
	}
}

func migrateSchema(db *gorm.DB) {
	db.AutoMigrate(NoteEntity{})
}

type NoteRepository interface {
	FindAll() []NoteEntity
	FindById(id int64) (NoteEntity, bool)
	Save(note NoteEntity) NoteEntity
	Delete(id int64)
}

func (it noteRepository) FindById(id int64) (NoteEntity, bool) {
	var note NoteEntity

	it.db.First(&note, id)
	if note.ID == 0 {
		return note, false
	}

	return note, true
}

func (it noteRepository) Save(note NoteEntity) NoteEntity {
	it.db.Save(&note)
	return note
}

func (it noteRepository) FindAll() []NoteEntity {
	var notes []NoteEntity

	it.db.Find(&notes)
	return notes
}

func (it noteRepository) Delete(id int64) {
	note, found := it.FindById(id)
	if !found {
		return
	}

	it.db.Delete(&note)
}
