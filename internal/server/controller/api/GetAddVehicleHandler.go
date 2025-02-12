package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/server/models/api"
)

// GetAddVehicleHandler returns a handler for the /api/vehicles route
func GetAddVehicleHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var Input api.VehicleInput
		if err := c.BodyParser(&Input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// TODO : Add the vehicle to the database

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "success"})
	}
}
