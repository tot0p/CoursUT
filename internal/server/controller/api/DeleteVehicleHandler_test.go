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

func TestDeleteVehicleHandler(t *testing.T) {
	var route = "/api/vehicles/:id"
	var routeValue = "/api/vehicles/1"
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
			description:   "Working request add vehicle",
			expectedError: false,
			expectedCode:  204,
			body:          "",
			expectedBody:  "",
		},
		{
			description:   "Not exist request Get vehicles",
			expectedError: false,
			body:          "",
			expectedCode:  404,
			expectedBody:  "{\"error\":\"Vehicle not found\"}",
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
	app := fiber.New()
	app.Delete(route, DeleteVehicleHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"DELETE",
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
	}
}
