package handler

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTicket(c *fiber.Ctx) error {
	// TODO: business logic
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Ticket has created"})
}

func GetTicket(c *fiber.Ctx) error {
	// TODO: business logic
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Ticket Found"})
}

func PurchaseTicket(c *fiber.Ctx) error {
	// TODO: business logic
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Purchase has created"})
}
