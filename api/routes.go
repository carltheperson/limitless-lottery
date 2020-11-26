package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carlriis/Limitless-Lottery/tickets"
)

func checkTicket(w http.ResponseWriter, r *http.Request) { // TODO create better validation
	id := r.URL.Query().Get("ticketid")
	amount := r.URL.Query().Get("amount")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' is a required field"))
		return
	}
	if _, err := strconv.Atoi(amount); err != nil || amount == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'amount' is not defined or a number"))
		return
	}

	ticketamount, _ := strconv.Atoi(amount)

	ct, err := tickets.Check(id, ticketamount)
	if err == tickets.ErrIDWithNoMatch {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' did not match any ticket"))
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}
