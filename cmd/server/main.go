package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/cmd/server/handler"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/domain"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/tickets"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/pkg/store"
)

func main() {

	// Cargo csv.
	/*list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	} else {
		fmt.Printf("Total tickets: %d", len(list))
	}*/

	db := store.New(store.FileType, "./tickets.csv")

	repository := tickets.NewRepository(db)
	service := tickets.NewService(repository)
	ticketHandler := handler.NewTickets(service)

	server := gin.Default()

	api := server.Group("/api/v1")

	api.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	t := api.Group("/tickets")
	{
		t.GET("/", ticketHandler.GetTickets())
		t.GET("/getByCountry/:dest", ticketHandler.GetTicketsByCountry())
		t.GET("/getPriceAverageByCountry/:dest", ticketHandler.GetPriceAverageByCountry())
	}

	if err := server.Run(":9000"); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
