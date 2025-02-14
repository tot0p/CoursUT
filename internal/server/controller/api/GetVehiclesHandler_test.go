package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"io"
	"net/http"
	"testing"
)

func TestGetVehiclesHandler(t *testing.T) {
	var route = "/api/vehicles"
	var routeValue = "/api/vehicles"
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
			description:   "Working request Get vehicles",
			expectedError: false,
			expectedCode:  204,
			body:          "",
			expectedBody:  "",
		},
		{
			description:   "Not exist request Get vehicles",
			expectedError: false,
			body:          "",
			expectedCode:  200,
			expectedBody:  "[{\"id\":1,\"plate\":\"AA-123-AA\",\"vehicleType\":1}]",
		},
	}
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get(route, GetVehiclesHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			routeValue,
			nil,
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
		vehicle.CreateVehicle(&models.Vehicle{
			Plate:       "AA-123-AA",
			VehicleType: models.Car,
		})
	}
}
