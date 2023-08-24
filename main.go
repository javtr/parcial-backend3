package main

import (
	"parcial/internal/tickets"
)

func main() {

	// Requerimiento 1
	total, err := tickets.GetTotalTickets("Brazil")

	println("El total de tickets con el pais solicitado es ", total)
	println(err)

	//Requerimiento 2
	tot2, err2 := tickets.GetCountByPeriod("21-22")

	println("La cantidad de tickets en ese rango son: ", tot2)
	println(err2)

}
