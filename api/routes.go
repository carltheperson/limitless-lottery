package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carlriis/Limitless-Lottery/tickets"
	"github.com/go-playground/validator/v10"
)

func checkTicket(w http.ResponseWriter, r *http.Request) {
	v := validator.New()
	input := struct {
		ID     string `validate:"required"`
		Amount string `validate:"required,numeric"`
		amount int
	}{
		ID:     r.URL.Query().Get("ticketid"),
		Amount: r.URL.Query().Get("amount"),
	}

	err := v.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(e.Error()))
			return
		}
	}

	num, _ := strconv.Atoi(input.Amount)
	input.amount = num

	if input.amount < 0 || input.amount > 100000000 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'amount' outside range"))
		return
	}

	ct, err := tickets.Check(input.ID, input.amount)
	if err == tickets.ErrIDWithNoMatch {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' did not match any ticket"))
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}
