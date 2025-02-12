package api

import "github.com/gofiber/fiber/v2"

// GetPingHandler returns a handler for the /api/ping route
func GetPingHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	}
}
