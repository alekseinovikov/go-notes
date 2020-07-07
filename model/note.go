package model

type Note struct {
	Id      int64
	Title   string
	Content string
}

type NoteCreateRequest struct {
	Title   string
	Content string
}
