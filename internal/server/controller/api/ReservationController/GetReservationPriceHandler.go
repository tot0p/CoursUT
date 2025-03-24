package ReservationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"strconv"
)

// GetReservationPriceHandler is the handler at /api/reservations/:id/price to get the price of a reservation
func GetReservationPriceHandler(c *fiber.Ctx) error {
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
	reservationVehicle, err := vehicle.GetVehicle(reservation.VehicleID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot get vehicle"})
	}
	// calculate the price
	price := models.GetDegresiveParkingPrice(reservationVehicle.VehicleType, reservation.ParkingDuration)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"price": price})
}
