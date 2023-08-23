package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID     string
	Nombre string
	Email  string
	Pais   string
	Hora   string
	Precio int
}

// funcion de lectura de archivo csv
func ReadTicketsFromCSV(filename string) ([]Ticket, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tickets []Ticket
	for _, line := range lines {
		if len(line) >= 6 {
			ticket := Ticket{
				ID:     line[0],
				Nombre: line[1],
				Email:  line[2],
				Pais:   line[3],
				Hora:   line[4],
				Precio: atoi(line[5]),
			}
			tickets = append(tickets, ticket)
		}
	}

	return tickets, nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

//Ej 1
func GetTotalTickets(destination string) (int, error) {
	tickets, err := ReadTicketsFromCSV("tickets.csv")
	if err != nil {
		return 0, err
	}

	count := 0
	for _, ticket := range tickets {
		if strings.TrimSpace(ticket.Pais) == destination {
			count++
		}
	}

	return count, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	tickets, err := ReadTicketsFromCSV("tickets.csv")
	if err != nil {
		return 0, err
	}

	startHour, _ := strconv.Atoi(strings.Split(time, "-")[0])
	endHour, _ := strconv.Atoi(strings.Split(time, "-")[1])

	count := 0
	for _, ticket := range tickets {
		ticketHour, _ := strconv.Atoi(strings.Split(ticket.Hora, ":")[0])
		if ticketHour >= startHour && ticketHour <= endHour {
			count++
		}
	}

	return count, nil
}



