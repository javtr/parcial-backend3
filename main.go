package main

import (
	"parcial/internal/tickets"
	"fmt"
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

	//Requerimiento 3
	tot3, err3 := tickets.AverageDestination("Brazil")
	if err3 != nil {
		fmt.Println("Error al calcular el promedio de tickets para Brazil:", err3)
		return
	}
	fmt.Println("El promedio de tickets a ese país es:", tot3, "%")


}
