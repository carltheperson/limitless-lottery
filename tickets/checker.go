package tickets

import (
	"errors"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// ErrIDWithNoMatch is thrown when the tickitid does not match any tickets
var ErrIDWithNoMatch = errors.New("ID did not match any tickets")

func checkTicket(ticket Ticket, amount int) CheckedTicket {
	checkedTicket := CheckedTicket{AmountWonTotal: 0, AmountDeducted: 0, Wins: []OddsWin{}}

	checkedTicket.AmountDeducted = ticket.price * amount

	for outOf, prize := range ticket.odds {

		// Using binomial distribution to determine how many of the tickets won using the 'outOf' odds.
		bio := distuv.Binomial{
			N:   float64(amount),
			P:   1.0 / float64(outOf),
			Src: rand.NewSource(uint64(time.Now().UnixNano())),
		}

		howManyWon := int(bio.Rand())

		if howManyWon != 0 {
			checkedTicket.AmountWonTotal += howManyWon * prize

			checkedTicket.Wins = append(checkedTicket.Wins, OddsWin{
				OutOfOdds:     outOf,
				Prize:         prize,
				AmountThatWon: howManyWon,
				TotalWinning:  howManyWon * prize,
			})
		}
	}

	return checkedTicket
}

// OddsWin represents the winnings of a certain set of odds. For example 1/10
type OddsWin struct {
	OutOfOdds     int // If the odds are 1/10. Then this number would be 10.
	Prize         int
	AmountThatWon int
	TotalWinning  int
}

// CheckedTicket represents a ticket after it has been checked
type CheckedTicket struct {
	AmountWonTotal int
	AmountDeducted int
	Wins           []OddsWin
}

// Check takes in a ticket ID and returns a CheckedTicket
func Check(id string, amount int) (CheckedTicket, error) {
	for _, ticket := range tickets {
		if ticket.id == id {
			return checkTicket(ticket, amount), nil
		}
	}
	return CheckedTicket{}, ErrIDWithNoMatch
}
