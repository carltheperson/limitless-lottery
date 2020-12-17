package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carlriis/Limitless-Lottery/api/validation"
	"github.com/carlriis/Limitless-Lottery/db"
	"github.com/carlriis/Limitless-Lottery/tickets"
)

var (
	errNoTicketMatch    = validation.ErrorMessage{Case: "match", Field: "ID", Message: "There was no match for that ticket id"}
	errUserDoesNotExist = validation.ErrorMessage{Case: "exist", Field: "Username", Message: "No user was found with that username"}
)

func checkTicketAmount(w http.ResponseWriter, r *http.Request) {
	amountInt, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	input := struct {
		ID       string `validate:"required"`
		Amount   int    `validate:"required,numeric,min=0,max=1000000000"`
		Username string `validate:"required"`
	}{
		ID:       r.URL.Query().Get("ticketid"),
		Amount:   amountInt,
		Username: r.URL.Query().Get("username"),
	}

	ea := validation.NewErrorAdder()
	validation.AddErrorsFromInput(input, &ea)
	if ea.HasErrors() == true {
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
	input := struct {
		ID       string `validate:"required"`
		Username string `validate:"required"`
	}{
		ID:       r.URL.Query().Get("ticketid"),
		Username: r.URL.Query().Get("username"),
	}

	ea := validation.NewErrorAdder()
	validation.AddErrorsFromInput(input, &ea)
	if ea.HasErrors() == true {
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
