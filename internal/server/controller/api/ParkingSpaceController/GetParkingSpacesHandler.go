package ParkingSpaceController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
)

// GetParkingSpacesHandler is the handler at /api/parking-spaces for get all parking spaces
func GetParkingSpacesHandler(c *fiber.Ctx) error {
	parkingSpaces, err := parkingSpace.GetParkingSpaces()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if len(parkingSpaces) == 0 {
		return c.Status(fiber.StatusNoContent).Send(nil)
	}
	return c.JSON(parkingSpaces)
}
