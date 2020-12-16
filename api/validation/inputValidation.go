package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func AddErrorsFromInput(input interface{}, ea *ErrorAdder) {
	reflected := reflect.ValueOf(input)

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
