package controllers

import (
	"github.com/akasaa101/ticketing/internal/models"
	"github.com/akasaa101/ticketing/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"strings"
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

// CreateTicket handles the creation of a new ticket
// @Summary Create a new ticket
// @Tags Tickets
// @Accept json
// @Produce json
// @Param ticket body models.Ticket true "Ticket Data"
// @Success 201 {object} map[string]interface{}
// @Router /tickets [post]
func (ts TicketService) CreateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket

	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid Json.",
			"message": err.Error(),
		})
	}

	insertedTicket, err := ts.Service.TicketInsert(ticket)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error":   "Service unavailable.",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id":         insertedTicket.Id,
		"name":       insertedTicket.Name,
		"desc":       insertedTicket.Desc,
		"allocation": insertedTicket.Allocation,
	})
}

// GetTicket godoc
// @Summary Get a ticket by ID
// @Description Retrieve a ticket's details by its ID
// @Tags Tickets
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} models.Ticket
// @Failure 400 {object} fiber.Map "Invalid ticket ID"
// @Failure 404 {object} fiber.Map "Ticket not found"
// @Router /tickets/{id} [get]
func (bs TicketService) GetTicket(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid ticket ID.",
			"message": err.Error(),
		})
	}

	ticket, err := bs.Service.TicketGetById(int16(id))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error":   "Ticket not found.",
				"message": err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Internal server error.",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":         ticket.ID,
		"name":       ticket.Name,
		"desc":       ticket.Desc,
		"allocation": ticket.Allocation,
	})
}

// PurchaseTicket godoc
// @Summary Purchase a ticket
// @Description Purchase a ticket by specifying the quantity
// @Tags Tickets
// @Accept json
// @Produce json
// @Param id path int true "Ticket ID"
// @Param quantity body int true "Quantity to purchase"
// @Success 200 {string} string "Purchase successful"
// @Failure 400 {object} fiber.Map "Invalid ticket ID or quantity"
// @Failure 404 {string} string "Ticket not found"
// @Failure 409 {string} string "Insufficient ticket allocation"
// @Router /tickets/{id}/purchases [post]
func (bs TicketService) PurchaseTicket(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return c.Status(http.StatusBadRequest).SendString("Invalid ticket ID.")
	}

	var request struct {
		Quantity int `json:"quantity"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to parse request body.")
	}

	err = bs.Service.PurchaseTicket(int16(id), request.Quantity)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(http.StatusNotFound).SendString("Ticket not found.")
		}
		if strings.Contains(err.Error(), "non-positive quantity") {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error":   "Invalid purchase quantity.",
				"message": err.Error(),
			})
		}
		if strings.Contains(err.Error(), "insufficient allocation") {
			return c.Status(http.StatusConflict).SendString("Insufficient ticket allocation.")
		}
		return c.Status(http.StatusInternalServerError).SendString("Internal server error.")
	}

	return c.SendStatus(http.StatusOK)
}
