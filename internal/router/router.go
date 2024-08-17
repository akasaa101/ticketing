package router

import (
	"github.com/akasaa101/ticketing/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	ticket := app.Group("/tickets")
	ticket.Post("/", handler.CreateTicket)
	ticket.Get("/:id", handler.GetTicket)
	ticket.Post("/:id/purchases", handler.PurchaseTicket)
}
