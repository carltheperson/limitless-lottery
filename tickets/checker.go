package tickets

import (
	"errors"
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// ErrIDWithNoMatch is thrown when the tickitid does not match any tickets
var ErrIDWithNoMatch = errors.New("ID did not match any tickets")

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

func calcAmountForOneWin(outOf int) int {
	randomNumber := rand.New(rand.NewSource(uint64(time.Now().UnixNano()))).Float64()
	amount := math.Ceil(math.Log(-1.0*randomNumber+1.0) / math.Log(1.0-1.0/float64(outOf)))
	return int(amount)
}

func checkTicketUntilWin(ticket Ticket) CheckedTicketUntilWin {
	minAmount := math.MaxInt32
	prizeForMinAmount := 0
	outOfForMinAmount := 0

	for outOf, prize := range ticket.odds {
		amount := calcAmountForOneWin(outOf)
		if amount < minAmount {
			minAmount = amount
			prizeForMinAmount = prize
			outOfForMinAmount = outOf
		}
	}

	return CheckedTicketUntilWin{
		Prize:               prizeForMinAmount,
		Price:               ticket.price,
		AmountBought:        minAmount,
		OutOfOdds:           outOfForMinAmount,
		AmountDeductedTotal: minAmount * ticket.price,
		Profit:              prizeForMinAmount - minAmount*ticket.price,
	}
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

type CheckedTicketUntilWin struct {
	Prize               int
	Price               int
	AmountBought        int
	OutOfOdds           int
	AmountDeductedTotal int
	Profit              int
}

func findCheckedTicketFromID(id string) (Ticket, error) {
	for _, ticket := range tickets {
		if ticket.id == id {
			return ticket, nil
		}
	}
	return Ticket{}, ErrIDWithNoMatch
}

// Check checks a ticket a certain amount
func CheckAmount(id string, amount int) (CheckedTicketAmount, error) {
	ticket, err := findCheckedTicketFromID(id)
	if err != nil {
		return CheckedTicketAmount{}, err
	}
	return checkTicketAmount(ticket, amount), nil
}

// CheckUntilWin buys tickets until a success
func CheckUntilWin(id string) (CheckedTicketUntilWin, error) {
	ticket, err := findCheckedTicketFromID(id)
	if err != nil {
		return CheckedTicketUntilWin{}, err
	}
	return checkTicketUntilWin(ticket), nil
}
