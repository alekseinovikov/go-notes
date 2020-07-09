package errors

type noteError struct {
	errorCode int
}

var NotFoundError = noteError{errorCode: 0}

type NoteError interface {
	GetError() int
	String() string
}

func (it noteError) String() string {
	return [...]string{"NotFoundError"}[it.errorCode]
}

func (it noteError) GetError() int {
	return it.errorCode
}
