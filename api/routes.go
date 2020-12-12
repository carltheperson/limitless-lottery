package api

import (
	"encoding/json"
	"net/http"

	"github.com/carlriis/Limitless-Lottery/api/validation"
	"github.com/carlriis/Limitless-Lottery/db"
	"github.com/carlriis/Limitless-Lottery/tickets"
)

var (
	noTicketMatch    = []validation.ErrorMessage{{Case: "match", Field: "ID", Message: "There was no match for that ticket id"}}
	userDoesNotExist = []validation.ErrorMessage{{Case: "Exist", Field: "Username", Message: "No user was found with that username"}}
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	input, ok := validation.CheckTicketAmount(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckAmount(input.ID, input.Amount)
	if err == tickets.ErrIDWithNoMatch {
		validation.WriteErrors(w, noTicketMatch)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(input.Username, ct.AmountWonTotal-ct.AmountDeducted)
	if err == db.ErrUserDoesNotExist {
		validation.WriteErrors(w, userDoesNotExist)
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
	input, ok := validation.CheckTicketUntilWin(w, r)
	if ok != true {
		return
	}

	ct, err := tickets.CheckUntilWin(input.ID)
	if err == tickets.ErrIDWithNoMatch {
		validation.WriteErrors(w, noTicketMatch)
		return
	}

	updatedBalance, err := db.ChangeUserBalance(input.Username, ct.Profit)
	if err == db.ErrUserDoesNotExist {
		validation.WriteErrors(w, userDoesNotExist)
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
