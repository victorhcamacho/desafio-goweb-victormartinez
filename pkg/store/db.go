package store

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/domain"
)

type StoreType string

type fileStore struct {
	FilePath string
}

const (
	FileType StoreType = "file"
)

type Store interface {
	Read() ([]domain.Ticket, error)
	Write([]domain.Ticket) error
}

func New(stroe StoreType, path string) Store {
	switch stroe {
	case FileType:
		return &fileStore{path}
	}
	return nil
}

func (fs *fileStore) Write(tickets []domain.Ticket) error {

	var lines string

	for _, ticket := range tickets {
		lines += fmt.Sprintf("%v,%v,%v,%v,%v,%.2f\n", ticket.Id, ticket.Name, ticket.Email, ticket.Country, ticket.Time, ticket.Price)
	}

	return os.WriteFile(fs.FilePath, []byte(lines), 0644)
}

func (fs *fileStore) Read() ([]domain.Ticket, error) {

	var tickets []domain.Ticket

	if content, errR := os.ReadFile(fs.FilePath); errR != nil {
		return nil, errR
	} else {

		lines := strings.Split(string(content), "\n")

		if len(lines) == 0 {
			return nil, fmt.Errorf("la lista de tickets esta vacia")
		}

		for _, line := range lines {
			if line != "" {

				pieces := strings.Split(line, ",")

				if price, errP := strconv.ParseFloat(pieces[5], 64); errP != nil {
					return nil, fmt.Errorf("no se pudo convertir el precio a formato flotante")
				} else {
					ticket := domain.NewTicket(pieces[0], pieces[1], pieces[2], pieces[3], pieces[4], price)
					tickets = append(tickets, *ticket)
				}
			}
		}

		return tickets, nil
	}
}
