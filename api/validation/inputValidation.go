package validation

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

// AddErrorsFromInput takes an input struct and adds validation errores using the ErrorAdder
func AddErrorsFromInput(input interface{}, ea *ErrorAdder) {
	var reflected reflect.Value
	if reflect.ValueOf(input).Kind() == reflect.Ptr {
		reflected = reflect.ValueOf(input).Elem()
	} else {
		reflected = reflect.ValueOf(input)
	}

	err := v.Struct(input)
	if err != nil {

		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			jsonField, _ := reflected.Type().FieldByName(err.StructField())

			field := jsonField.Tag.Get("json")
			if field == "" {
				field = strings.ToLower(err.StructField())
			}

			errorMessage := createErrorMessage(field, err.Tag())
			ea.Add(errorMessage)
		}
	}
}

// UnmarshalJSONAndAddErrors turns the JSON into a struct and adds relevant validation errors, including syntax errors
func UnmarshalJSONAndAddErrors(input interface{}, body io.ReadCloser, ea *ErrorAdder) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	jsonString := buf.String()

	err := json.Unmarshal([]byte(jsonString), &input)
	if _, ok := err.(*json.SyntaxError); ok {
		ea.Add(ErrorMessage{Case: "syntaxError", Field: "JSON", Message: "SyntaxError parsing JSON"})
		return
	} else if v, ok := err.(*json.UnmarshalTypeError); ok {
		ea.Add(ErrorMessage{Case: "JSONTypeError", Field: v.Field, Message: "Field " + v.Field + " cannot be " + v.Value})
		return
	} else if err != nil {
		ea.Add(ErrorMessage{Case: "unexpectedError", Field: "JSON", Message: "The server had an unexpected error while trying to parse JSON"})
		return
	}

	AddErrorsFromInput(input, ea)
}
