package tickets

import (
	"context"

	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/domain"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/pkg/store"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if tickets, errReading := r.db.Read(); errReading != nil {
		return []domain.Ticket{}, errReading
	} else {
		return tickets, nil
	}
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	if tickets, errReading := r.db.Read(); errReading != nil {
		return []domain.Ticket{}, errReading
	} else {
		var ticketsDest []domain.Ticket
		for _, t := range tickets {
			if t.Country == destination {
				ticketsDest = append(ticketsDest, t)
			}
		}
		return ticketsDest, nil
	}
}
