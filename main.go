package main

import (
	"github.com/carlriis/Limitless-Lottery/api"
	"github.com/carlriis/Limitless-Lottery/db"
)

func main() {
	db.Connect()
	api.Serve(":8080")
}
