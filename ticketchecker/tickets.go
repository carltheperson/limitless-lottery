package ticketchecker

type ticket struct {
	id           string
	price        int
	calculateWin func() int
}

var scratchy = ticket{id: "scr", price: 10, calculateWin: func() int {
	if checkWin(10) {
		return 10
	}
	if checkWin(15) {
		return 20
	}
	if checkWin(25) {
		return 100
	}
	if checkWin(35) {
		return 200
	}
	if checkWin(55) {
		return 300
	}
	if checkWin(150) {
		return 1000
	}
	return 0
}}

var goldenTicket = ticket{id: "gol", price: 20, calculateWin: func() int {
	if checkWin(100) {
		return 500
	}
	if checkWin(150) {
		return 1000
	}
	if checkWin(200) {
		return 1500
	}
	if checkWin(250) {
		return 2000
	}
	return 0
}}

var insaneMoneyRain = ticket{id: "ins", price: 30, calculateWin: func() int {
	if checkWin(100000) {
		return 10000000
	}
	return 0
}}

var tickets = []ticket{scratchy, goldenTicket, insaneMoneyRain}
