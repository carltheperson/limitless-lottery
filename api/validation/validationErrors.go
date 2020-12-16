package validation

// ErrorMessage represents one error to be returned to the user
type ErrorMessage struct {
	Case    string
	Field   string
	Message string
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
