package ReservationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"strconv"
)

// GetReservationHandler is the handler at /api/reservations to get a reservation
func GetReservationHandler(c *fiber.Ctx) error {
	// get the reservation ID
	reservationID := c.Params("id")
	realId, err := strconv.Atoi(reservationID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// get the reservation
	reservation, err := parkingSpaceInformation.GetParkingSpaceInformation(realId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot get reservation"})
	}
	return c.Status(fiber.StatusOK).JSON(reservation)
}
