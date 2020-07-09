package controller

import (
	"github.com/alekseinovikov/go-notes/errors"
	"net/http"
)

func HandleError(err errors.NoteError, w http.ResponseWriter) {
	var statusCode int
	switch err {
	case errors.NotFoundError:
		statusCode = http.StatusNotFound
	}

	http.Error(w, err.String(), statusCode)
}
