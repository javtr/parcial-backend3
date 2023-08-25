package main

import (
	"fmt"
	"parcial/internal/tickets"
)

func main() {
	totalTicketsChannel := make(chan int)
	morningsChannel := make(chan int)
	averageDestinationChannel := make(chan int)

	// Requerimiento 1: Obtener el total de tickets para "Brazil"
	go func() {
		totalTickets, err := tickets.GetTotalTickets("Brazzczil")
		if err != nil {
			fmt.Println("Error en la obtención del total de tickets para Brazil:", err)
			panic(err)
		}
		totalTicketsChannel <- totalTickets // Enviar el resultado al canal
	}()

	// Requerimiento 2: Obtener la cantidad de tickets en el rango horario "earlyMorning"
	go func() {
		earlyMorning, err := tickets.GetCountByPeriod("0-6")
		if err != nil {
			fmt.Println("Error al obtener la cantidad de tickets en ese rango:", err)
			panic(err)
		}
		morningsChannel <- earlyMorning // Enviar el resultado al canal
	}()

	// Requerimiento 3: Calcular el promedio de tickets para "Brazil"
	go func() {
		averageDestination, err := tickets.AverageDestination("Brazil")
		if err != nil {
			fmt.Println("Error al calcular el promedio de tickets para Brazil:", err)
			panic(err)
		}
		averageDestinationChannel <- averageDestination // Enviar el resultado al canal
	}()

	// Recibir los resultados de los canales
	totalTicketsResult := <-totalTicketsChannel
	morningsResult := <-morningsChannel
	averageDestinationResult := <-averageDestinationChannel

	fmt.Println("El total de tickets con el país solicitado es", totalTicketsResult)
	fmt.Println("La cantidad de tickets en ese rango es:", morningsResult)
	fmt.Println("El promedio de tickets a ese país es:", averageDestinationResult, "%")
}
