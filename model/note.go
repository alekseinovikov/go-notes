package model

type Note struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
