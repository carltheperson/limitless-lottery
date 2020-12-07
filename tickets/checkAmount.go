package tickets

import (
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

func checkTicketAmount(ticket Ticket, amount int) CheckedTicketAmount {
	checkedTicket := CheckedTicketAmount{AmountWonTotal: 0, AmountDeducted: 0, Wins: []OddsWin{}}

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

// CheckedTicketAmount represents a ticket after it has been checked
// Note that tickets can have multiple odds with different winning prizes
type CheckedTicketAmount struct {
	AmountWonTotal int
	AmountDeducted int
	Wins           []OddsWin
}

// CheckAmount checks a ticket a certain amount
func CheckAmount(id string, amount int) (CheckedTicketAmount, error) {
	ticket, err := findCheckedTicketFromID(id)
	if err != nil {
		return CheckedTicketAmount{}, err
	}
	return checkTicketAmount(ticket, amount), nil
}
