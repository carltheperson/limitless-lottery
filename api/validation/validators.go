package validation

import (
	"net/http"
	"strconv"
)

type CheckTicketAmountInput struct {
	ID       string `validate:"required"`
	Amount   int    `validate:"required,numeric,min=0,max=1000000000"`
	Username string `validate:"required"`
}

func CheckTicketAmount(w http.ResponseWriter, r *http.Request) (CheckTicketAmountInput, bool) {
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))

	input := CheckTicketAmountInput{
		ID:       r.URL.Query().Get("ticketid"),
		Amount:   amount,
		Username: r.URL.Query().Get("username"),
	}

	ok := validateInput(input, w)
	if ok == false {
		return CheckTicketAmountInput{}, false
	}

	return input, true
}

type CheckTicketUntilWinInput struct {
	ID       string `validate:"required"`
	Username string `validate:"required"`
}

func CheckTicketUntilWin(w http.ResponseWriter, r *http.Request) (CheckTicketUntilWinInput, bool) {
	input := CheckTicketUntilWinInput{
		ID:       r.URL.Query().Get("ticketid"),
		Username: r.URL.Query().Get("username"),
	}

	ok := validateInput(input, w)
	if ok == false {
		return CheckTicketUntilWinInput{}, false
	}
	return input, true
}
