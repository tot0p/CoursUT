package ReservationController

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"io"
	"net/http"
	"testing"
)

func TestAddReservationHandler(t *testing.T) {
	var route = "/api/reservations"
	tests := []struct {
		description string

		// Test input
		body string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Bad request add reservation wrong duration",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"vehicle_id\":1,\"reservation_time\":\"toto\"}",
			expectedBody:  "{\"error\":\"Cannot parse duration\"}",
		},
		{
			description:   "Working request add reservation",
			expectedError: false,
			body:          "{\"vehicle_id\":1,\"reservation_time\":\"1h\"}",
			expectedCode:  201,
			expectedBody:  "{\"id\":1,\"vehicle_id\":1,\"parking_space_id\":1,\"arrival_time\":\"0001-01-01T00:00:00Z\",\"departure_time\":\"0001-01-01T00:00:00Z\",\"parking_duration\":3600000000000}",
		},
		{
			description:   "Bad request add reservation",
			expectedError: false,
			expectedCode:  400,
			body:          "",
			expectedBody:  "{\"error\":\"Cannot parse input\"}",
		},
		{
			description:   "Bad request add reservation",
			expectedError: false,
			expectedCode:  400,
			body:          "{}",
			expectedBody:  "{\"error\":\"Cannot get vehicle\"}",
		},
		{
			description:   "Bad request add reservation wrong vehicle type",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"vehicle_id\":2,\"reservation_time\":\"1h\"}",
			expectedBody:  "{\"error\":\"No available parking space\"}",
		},
		{
			description:   "Bad request no available parking space",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"vehicle_id\":1,\"reservation_time\":\"1h\"}",
			expectedBody:  "{\"error\":\"No available parking space\"}",
		},
	}

	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	_, err = vehicle.CreateVehicle(&models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}
	_, err = vehicle.CreateVehicle(&models.Vehicle{
		Plate:       "BB-123-BB",
		VehicleType: models.Truck,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Post(route, AddReservationHandler)
	for _, test := range tests {
		var bodyR io.Reader = nil
		if test.body != "" {
			bodyR = bytes.NewBuffer([]byte(test.body))
		}
		req, _ := http.NewRequest(
			"POST",
			route,
			bodyR,
		)
		req.Header.Set("Content-Type", "application/json")
		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		body, err := io.ReadAll(res.Body)
		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
