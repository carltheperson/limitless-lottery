package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/api/validation"
	"github.com/carlriis/Limitless-Lottery/tickets"
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	input, ok := validation.CheckTicketAmount(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckAmount(input.ID, input.Amount)
	if err == tickets.ErrIDWithNoMatch {
		validation.WriteErrors(w, []validation.ErrorMessage{{Case: "match", Field: "ID"}})
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}

func checkTicketUntilWin(w http.ResponseWriter, r *http.Request) {
	input, ok := validation.CheckTicketUntilWin(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckUntilWin(input.ID)
	if err == tickets.ErrIDWithNoMatch {
		validation.WriteErrors(w, []validation.ErrorMessage{{Case: "match", Field: "ID"}})
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}
