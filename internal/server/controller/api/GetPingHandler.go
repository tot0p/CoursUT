package api

import "github.com/gofiber/fiber/v2"

func GetPingHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	}
}
