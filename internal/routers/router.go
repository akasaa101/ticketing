package routers

import (
	"github.com/akasaa101/ticketing/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, tc controllers.TicketController) {
	ticket := app.Group("/tickets")
	ticket.Post("/", tc.CreateTicket)
	ticket.Get("/:id", tc.GetTicket)
	ticket.Post("/:id/purchases", tc.PurchaseTicket)
}
