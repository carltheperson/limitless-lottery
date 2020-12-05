package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/tickets"
)

func checkTicket(w http.ResponseWriter, r *http.Request) {
	input, ok := checkTicketValidator(w, r)
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
