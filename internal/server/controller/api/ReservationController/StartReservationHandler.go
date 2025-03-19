package ReservationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"net/http"
	"strconv"
	"time"
)

func StartReservationHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	reservation, err := parkingSpaceInformation.GetParkingSpaceInformation(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}

	reservation.ArrivalTime = time.Now().UTC()
	err = parkingSpaceInformation.UpdateParkingSpaceInformation(reservation)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start reservation"})
	}

	return c.Status(http.StatusOK).JSON(reservation)
}
