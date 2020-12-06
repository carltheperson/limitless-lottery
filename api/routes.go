package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/tickets"
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	input, ok := checkTicketAmountValidator(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckAmount(input.ID, input.amount)
	if err == tickets.ErrIDWithNoMatch {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' did not match any ticket"))
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}

func checkTicketUntilWin(w http.ResponseWriter, r *http.Request) {
	input, ok := checkTicketUntilWinValidator(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckUntilWin(input.ID)
	if err == tickets.ErrIDWithNoMatch {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' did not match any ticket"))
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}
