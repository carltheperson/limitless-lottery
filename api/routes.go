package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/api/validation"
	"github.com/carlriis/Limitless-Lottery/db"
	"github.com/carlriis/Limitless-Lottery/tickets"
)

var (
	errNoTicketMatch    = validation.ErrorMessage{Case: "match", Field: "ID", Message: "There was no match for that ticket id"}
	errUserDoesNotExist = validation.ErrorMessage{Case: "exist", Field: "Username", Message: "No user was found with that username"}
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()
	input, ok := validation.CheckTicketAmount(r, &ea)
	if ok != true {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	ct, err := tickets.CheckAmount(input.ID, input.Amount)

	if err == tickets.ErrIDWithNoMatch {
		ea.Add(errNoTicketMatch)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(input.Username, ct.AmountWonTotal-ct.AmountDeducted)

	if err == db.ErrUserDoesNotExist {
		ea.Add(errUserDoesNotExist)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Ct      tickets.CheckedTicketAmount
		Balance int
	}{
		Ct:      ct,
		Balance: updatedBalance,
	})
}

func checkTicketUntilWin(w http.ResponseWriter, r *http.Request) {
	ea := validation.NewErrorAdder()
	input, ok := validation.CheckTicketUntilWin(r, &ea)
	if ok != true {
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	ct, err := tickets.CheckUntilWin(input.ID)

	if err == tickets.ErrIDWithNoMatch {
		ea.Add(errNoTicketMatch)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(input.Username, ct.Profit)

	if err == db.ErrUserDoesNotExist {
		ea.Add(errUserDoesNotExist)
		ea.Flush(w, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Ct      tickets.CheckedTicketUntilWin
		Balance int
	}{
		Ct:      ct,
		Balance: updatedBalance,
	})
}
