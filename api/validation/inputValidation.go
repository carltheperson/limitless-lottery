package validation

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func validateInput(input interface{}, w http.ResponseWriter) bool {
	reflected := reflect.ValueOf(input)
	errors := []ErrorMessage{}

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

			addError(field, err.Tag(), &errors)

		}
		WriteErrors(w, errors)
		return false
	}
	return true
}
