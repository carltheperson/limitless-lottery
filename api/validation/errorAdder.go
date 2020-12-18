package validation

import (
	"encoding/json"
	"net/http"
)

// ErrorAdder is a struct used to add and write relevant request errors
type ErrorAdder struct {
	Errors *[]ErrorMessage
}

// NewErrorAdder creates a new ErrorAdder
func NewErrorAdder() ErrorAdder {
	return ErrorAdder{Errors: &[]ErrorMessage{}}
}

// Add adds a new ErrorMessage to the ErrorAdder
func (ea ErrorAdder) Add(errorMessage ErrorMessage) {
	*ea.Errors = append(*ea.Errors, errorMessage)
}

// Flush writes all the current errors in the ErrorAdder, and sets a relevant http status
func (ea ErrorAdder) Flush(w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
	errorResponse := struct {
		Errors []ErrorMessage
	}{
		Errors: *ea.Errors,
	}
	json.NewEncoder(w).Encode(errorResponse)
}

// HasErrors checks whether there are any errors in the ErrorAdder
func (ea ErrorAdder) HasErrors() bool {
	return len(*ea.Errors) != 0
}
