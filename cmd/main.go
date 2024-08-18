package main

import (
	"github.com/akasaa101/ticketing/internal/controllers"
	"github.com/akasaa101/ticketing/internal/database"
	"github.com/akasaa101/ticketing/internal/repositories"
	"github.com/akasaa101/ticketing/internal/routers"
	"github.com/akasaa101/ticketing/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	ticketRepoDB := repositories.NewTicketRepositoryDB()
	ticketService := services.NewTicketService(ticketRepoDB)
	ticketController := controllers.NewTicketController(ticketService)

	routers.SetupRoutes(app, ticketController)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
