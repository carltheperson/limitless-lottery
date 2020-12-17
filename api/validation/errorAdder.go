package validation

import (
	"encoding/json"
	"net/http"
)

type ErrorAdder struct {
	Errors *[]ErrorMessage
}

func NewErrorAdder() ErrorAdder {
	return ErrorAdder{Errors: &[]ErrorMessage{}}
}

func (ea ErrorAdder) Add(errorMessage ErrorMessage) {
	*ea.Errors = append(*ea.Errors, errorMessage)
}

func (ea ErrorAdder) Flush(w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
	errorResponse := struct {
		Errors []ErrorMessage
	}{
		Errors: *ea.Errors,
	}
	json.NewEncoder(w).Encode(errorResponse)
}

func (ea ErrorAdder) HasErrors() bool {
	return len(*ea.Errors) != 0
}
