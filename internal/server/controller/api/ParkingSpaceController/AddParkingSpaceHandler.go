package ParkingSpaceController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/models"
)

func AddParkingSpaceHandler(c *fiber.Ctx) error {
	// Body parsing
	var Input models.ParkingSpace
	if err := c.BodyParser(&Input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse input"})
	}

	// Check if the parking space is valid
	if !models.IsValidParkingSpace(Input.SpaceNumber) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "SpaceNumber is not valid (format: letter followed by a number (000 - 999))"})
	}

	// Add parking space
	space, err := parkingSpace.CreateParkingSpace(&Input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(space)
}
