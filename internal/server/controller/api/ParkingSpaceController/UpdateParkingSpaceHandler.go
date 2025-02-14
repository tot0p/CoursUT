package ParkingSpaceController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/models"
)

func UpdateParkingSpaceHandler(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is not a number"})
	}
	// get body
	var input models.ParkingSpace
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse input"})
	}
	// check if the parking space is valid
	if !models.IsValidParkingSpace(input.SpaceNumber) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "SpaceNumber is not valid (format: letter followed by a number (000 - 999))"})
	}
	// update parking space
	err = parkingSpace.UpdateParkingSpace(&models.ParkingSpace{
		ID:          id,
		SpaceNumber: input.SpaceNumber,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
