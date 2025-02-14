package ParkingSpaceController

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
)

// DeleteParkingSpacesHandler is the handler at /api/parking-spaces/:id for delete a parking space
func DeleteParkingSpacesHandler(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is not a number"})
	}

	// check if the parking space exists
	_, err = parkingSpace.GetParkingSpace(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Parking space not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = parkingSpace.DeleteParkingSpace(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
