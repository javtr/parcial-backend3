package main

import (
	"parcial/internal/tickets"
)

func main() {

	// Requerimiento 1
	total, err := tickets.GetTotalTickets("Brazil")

	println("El total de tickets con el pais solicitado es ", total)
	println(err)

}
