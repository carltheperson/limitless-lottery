package api

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

func validate(input interface{}) validator.FieldError {
	err := v.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return e
		}
	}
	return nil
}

type checkTicketInput struct {
	ID     string `validate:"required"`
	Amount string `validate:"required,numeric"`
	amount int
}

func checkTicketValidator(w http.ResponseWriter, r *http.Request) (checkTicketInput, bool) {
	input := checkTicketInput{
		ID:     r.URL.Query().Get("ticketid"),
		Amount: r.URL.Query().Get("amount"),
	}

	err := validate(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return checkTicketInput{}, false
	}

	if input.amount < 0 || input.amount > 100000000 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'amount' outside range"))
		return checkTicketInput{}, false
	}

	num, _ := strconv.Atoi(input.Amount)
	input.amount = num

	return input, true
}
