package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/domain"
)

type Service interface {
	GetTickets(ctx gin.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx gin.Context, destination string) ([]domain.Ticket, error)
	GetPriceAverage(ctx gin.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTickets(ctx gin.Context) ([]domain.Ticket, error) {
	return s.repository.GetAll(&ctx)
}

func (s *service) GetTotalTickets(ctx gin.Context, destination string) ([]domain.Ticket, error) {
	return s.repository.GetTicketByDestination(&ctx, destination)
}

func (s *service) GetPriceAverage(ctx gin.Context, destination string) (float64, error) {

	tickets, err := s.repository.GetTicketByDestination(&ctx, destination)

	if err != nil || len(tickets) == 0 {
		return float64(0), err
	}

	var totalPrices float64
	for _, t := range tickets {
		totalPrices += t.Price
	}

	return totalPrices / float64(len(tickets)), nil
}
