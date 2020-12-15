package validation

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage represents one error to be returned to the user
type ErrorMessage struct {
	Case    string
	Field   string
	Message string
}

type ErrorAdder struct {
	Errors *[]ErrorMessage
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

func NewErrorAdder() ErrorAdder {
	return ErrorAdder{Errors: &[]ErrorMessage{}}
}

func createErrorMessage(field string, errorCase string) ErrorMessage {
	msg := generateHumanMessage(field, errorCase)
	return ErrorMessage{Case: errorCase, Field: field, Message: msg}
}

func generateHumanMessage(field string, errorCase string) string {
	var msg string

	switch errorCase {

	case "required":
		msg = "The " + field + " is required"

	case "numeric":
		msg = "The " + field + " is not numeric"

	case "min":
		msg = field + " is too low"

	case "max":
		msg = field + " is too high"

	default:
		msg = field + " is invalid"

	}

	return msg
}
