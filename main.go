package main

import (
	"fmt"
	"parcial/internal/tickets"
)

func main() {

	tickets, err := tickets.ReadTicketsFromCSV("tickets.csv")

	println(err)

	for _, elemento := range tickets {
		fmt.Println(elemento)
	}

}
