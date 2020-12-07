package tickets

import "errors"

// ErrIDWithNoMatch is thrown when the tickitid does not match any tickets
var ErrIDWithNoMatch = errors.New("ID did not match any tickets")

func findCheckedTicketFromID(id string) (Ticket, error) {
	for _, ticket := range tickets {
		if ticket.id == id {
			return ticket, nil
		}
	}
	return Ticket{}, ErrIDWithNoMatch
}

// Ticket represents a lottery ticket
type Ticket struct {
	id    string
	price int
	odds  map[int]int
}

var scratchy = Ticket{id: "scr", price: 10, odds: map[int]int{
	10:  10,
	15:  20,
	25:  100,
	35:  200,
	55:  300,
	150: 1000,
}}

var goldenTicket = Ticket{id: "gol", price: 20, odds: map[int]int{
	100: 500,
	150: 1000,
	200: 1500,
	250: 2000,
}}

var insaneMoneyRain = Ticket{id: "ins", price: 30, odds: map[int]int{
	100000: 10000000,
}}

var tickets = []Ticket{scratchy, goldenTicket, insaneMoneyRain}
