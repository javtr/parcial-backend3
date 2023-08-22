package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
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
