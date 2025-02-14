package VehicleController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"github.com/tot0p/CoursUT/internal/models/api"
	"github.com/tot0p/CoursUT/internal/utils"
)

// UpdateVehicleHandler is the handler for the update vehicle route
func UpdateVehicleHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Id is not valid",
		})
	}

	// get body
	var Input api.VehicleInput
	err = c.BodyParser(&Input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse input",
		})
	}

	//check if vehicle is valid
	if Input.Plate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Plate is required",
		})
	} else if Input.Type == models.Unknown || !models.IsValidVehicleType(int(Input.Type)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Type is required and must be different of Unknown"})
	} else if !utils.CheckPlate(Input.Plate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Plate is not valid"})
	}
	// update vehicle
	err = vehicle.UpdateVehicle(&models.Vehicle{
		ID:          id,
		Plate:       Input.Plate,
		VehicleType: Input.Type,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
