package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/tickets"
)

func checkTicket(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("ticketid")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' is a required field"))
		return
	}

	ct, err := tickets.Check(id)
	if err == tickets.IDWithNoMatchError {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'ticketid' did not match any ticket"))
		return
	}

	json.NewEncoder(w).Encode(ct)
	return
}
