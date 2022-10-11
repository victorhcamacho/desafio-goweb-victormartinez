package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/internal/tickets"
	"github.com/victorhcamacho/desafio-goweb-victormartinez/pkg/web"
)

type ticketHandler struct {
	service tickets.Service
}

func NewTickets(s tickets.Service) *ticketHandler {
	return &ticketHandler{
		service: s,
	}
}

func (s *ticketHandler) GetTickets() gin.HandlerFunc {
	return func(c *gin.Context) {

		tickets, err := s.service.GetTickets(*c)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				web.NewResponse(500, nil, err.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, tickets, ""))
	}
}

func (s *ticketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(*c, destination)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				web.NewResponse(500, nil, err.Error()),
			)
			return
		} else if len(tickets) == 0 {
			c.JSON(
				http.StatusNotFound,
				web.NewResponse(404, nil, fmt.Sprintf("No se encontraron tickets con destino a '%v'", destination)),
			)
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, tickets, ""))
	}
}

func (s *ticketHandler) GetPriceAverageByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.GetPriceAverage(*c, destination)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				web.NewResponse(500, nil, err.Error()),
			)
			return
		} else if avg == 0 {
			c.JSON(
				http.StatusNotFound,
				web.NewResponse(404, nil, fmt.Sprintf("No se encontraron tickets con destino a '%v'", destination)),
			)
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, avg, ""))
	}
}
