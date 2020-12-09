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

// WriteErrors writes the error messages to the writer
func WriteErrors(w http.ResponseWriter, errors []ErrorMessage) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errors)
}

type errorAdder struct {
	errors    *[]ErrorMessage
	errorCase string
	field     string
}

func (ea errorAdder) add(message string) {
	*ea.errors = append(*ea.errors, ErrorMessage{Case: ea.errorCase, Field: ea.field, Message: message})
}

func addError(field string, errorCase string, errors *[]ErrorMessage) {
	ea := errorAdder{errors: errors, errorCase: errorCase, field: field}

	switch errorCase {

	case "required":
		ea.add("The " + field + " is required")

	case "numeric":
		ea.add("The " + field + " is not numeric")

	case "min":
		ea.add(field + " is too low")

	case "max":
		ea.add(field + " is too high")

	default:
		ea.add(field + " is invalid")

	}
}
