package VehicleController

import "github.com/gofiber/fiber/v2"

// PingHandler is handler at /api/ping responds with a pong message
func PingHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "pong"})
}
