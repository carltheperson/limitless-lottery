package tickets

import (
	"math"
	"time"

	"golang.org/x/exp/rand"
)

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

// CheckedTicketUntilWin is what is returned to the user after they choose to buy until win
type CheckedTicketUntilWin struct {
	Prize               int
	Price               int
	AmountBought        int
	OutOfOdds           int
	AmountDeductedTotal int
	Profit              int
}

// CheckUntilWin buys tickets until a success
func CheckUntilWin(id string) (CheckedTicketUntilWin, error) {
	ticket, err := findCheckedTicketFromID(id)
	if err != nil {
		return CheckedTicketUntilWin{}, err
	}
	return checkTicketUntilWin(ticket), nil
}
