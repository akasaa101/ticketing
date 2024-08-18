package controllers

import (
	"github.com/akasaa101/ticketing/internal/models"
	"github.com/akasaa101/ticketing/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type TicketService struct {
	Service services.TicketService
}

type TicketController interface {
	CreateTicket(c *fiber.Ctx) error
	GetTicket(c *fiber.Ctx) error
	PurchaseTicket(c *fiber.Ctx) error
}

func NewTicketController(service services.TicketService) *TicketService {
	return &TicketService{service}
}

func (ts TicketService) CreateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket

	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid Json.",
			"message": err.Error(),
		})
	}

	err := ts.Service.TicketInsert(ticket)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error":   "Service unavailable.",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Ticket created successfully.",
	})
}

func (bs TicketService) GetTicket(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	result, err := bs.Service.TicketGetById(int16(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (ts TicketService) PurchaseTicket(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
