package ReservationController

import (
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetRemainingTime calculates the remaining time for a reservation
func GetRemainingTimeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	reservationID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reservation ID",
		})
	}

	reservation, err := parkingSpaceInformation.GetParkingSpaceInformation(reservationID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Reservation not found",
		})
	}

	if reservation.ArrivalTime.IsZero() {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Reservation not started",
		})
	}

	remainingTime := reservation.ArrivalTime.Add(reservation.ParkingDuration).Sub(time.Now().UTC())
	return c.JSON(fiber.Map{
		"remaining_time": remainingTime.Truncate(time.Second).String(),
	})
}
