package ticketchecker

import (
	"errors"
	"math/rand"
	"time"
)

func checkWin(odds int) bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(odds) == 1 {
		return true
	}
	return false
}

// CheckedTicket represents a ticket after it has been checked
type CheckedTicket struct {
	amountWon, amountDeducted int
}

// Check takes in a ticket ID and returns a CheckedTicket
func Check(id string) (CheckedTicket, error) {
	for _, ticket := range tickets {
		if ticket.id == id {
			return CheckedTicket{amountWon: ticket.calculateWin(), amountDeducted: ticket.price}, nil
		}
	}
	return CheckedTicket{}, errors.New("ID did not match any tickets")
}
