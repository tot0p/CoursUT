package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
)

func GetVehiclesHandler(c *fiber.Ctx) error {
	vehicles, err := vehicle.GetVehicles()
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if vehicles == nil || len(vehicles) == 0 {
		return c.Status(fiber.StatusNoContent).Send(nil)
	}
	return c.JSON(vehicles)
}
