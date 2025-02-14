package ParkingSpaceController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/models"
)

func AddParkingSpaceHandler(c *fiber.Ctx) error {

	// Body parsing
	var Input models.ParkingSpace
	if err := c.BodyParser(&Input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse input"})
	}

	// Check if the parking space is valid

	return nil
}
