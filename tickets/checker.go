package tickets

import (
	"errors"
	"math/rand"
	"time"
)

// IDWithNoMatchError is thrown when the tickitid does not match any tickets
var IDWithNoMatchError = errors.New("ID did not match any tickets")

func checkWin(odds int) bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(odds) == 1 {
		return true
	}
	return false
}

// CheckedTicket represents a ticket after it has been checked
type CheckedTicket struct {
	AmountWon      int `json:"amountWon"`
	AmountDeducted int `json:"amountDeducted"`
}

// Check takes in a ticket ID and returns a CheckedTicket
func Check(id string) (CheckedTicket, error) {
	for _, ticket := range tickets {
		if ticket.id == id {
			return CheckedTicket{AmountWon: ticket.calculateWin(), AmountDeducted: ticket.price}, nil
		}
	}
	return CheckedTicket{}, IDWithNoMatchError
}
