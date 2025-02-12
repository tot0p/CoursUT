package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"github.com/tot0p/CoursUT/internal/models/api"
	"github.com/tot0p/CoursUT/internal/utils"
)

// AddVehicleHandler is the handler at /api/vehicles for add a vehicle
func AddVehicleHandler(c *fiber.Ctx) error {
	var Input api.VehicleInput
	if err := c.BodyParser(&Input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse input"})
	}
	if Input.Plate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Plate is required"})
	} else if Input.Type == models.Unknown || !models.IsValidVehicleType(int(Input.Type)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Type is required and must be different of Unknown"})
	} else if !utils.CheckPlate(Input.Plate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Plate is not valid"})
	}

	v, err := vehicle.CreateVehicle(&models.Vehicle{
		Plate:       Input.Plate,
		VehicleType: Input.Type,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(v)
}
