package ReservationController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"github.com/tot0p/CoursUT/internal/models/api"
	"time"
)

// AddReservationHandler is the handler at /api/parking-spaces for add a parking space
func AddReservationHandler(c *fiber.Ctx) error {
	// Body parsing
	var Input api.ReservationInput
	var spaceInformation models.ParkingSpaceInformation
	if err := c.BodyParser(&Input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse input"})
	}
	vechicle, err := vehicle.GetVehicle(Input.VehicleID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot get vehicle"})
	}
	// get the first available parking space for that vehicle type
	space, err := parkingSpace.GetAvailableParkingSpace(vechicle.VehicleType)
	if err != nil {
		// no available parking space
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No available parking space"})
	}
	// create the reservation
	spaceInformation.ParkingSpaceID = space.ID
	spaceInformation.VehicleID = vechicle.ID
	spaceInformation.ParkingDuration, err = time.ParseDuration(Input.ReservationDuration)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse duration"})
	}
	temp, err := parkingSpaceInformation.CreateParkingSpaceInformation(&spaceInformation)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot create reservation"})
	}
	spaceInformation = *temp

	return c.Status(fiber.StatusCreated).JSON(spaceInformation)
}
