package main

import (
	"parcial/internal/tickets"
)

func main() {

	// tickets, err := tickets.ReadTicketsFromCSV("tickets.csv")

	// println(err)

	// for _, elemento := range tickets {
	// 	fmt.Println(elemento)
	// }

	total, err := tickets.GetTotalTickets("Brazil")

	println("El total de tickets con el pais solicitado es ", total)
	println(err)

}
