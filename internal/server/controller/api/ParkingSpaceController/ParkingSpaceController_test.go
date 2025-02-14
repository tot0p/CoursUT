package ParkingSpaceController

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/models"
	"io"
	"net/http"
	"testing"
)

func TestGetParkingSpacesHandler(t *testing.T) {
	var route = "/api/parking-spaces"
	var routeValue = "/api/parking-spaces"
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
			description:   "Working request Get parking-spaces",
			expectedError: false,
			expectedCode:  204,
			body:          "",
			expectedBody:  "",
		},
		{
			description:   "Get parking-spaces",
			expectedError: false,
			body:          "",
			expectedCode:  200,
			expectedBody:  "[{\"id\":1,\"space_number\":\"A001\"}]",
		},
	}
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get(route, GetParkingSpacesHandler)
	for i, test := range tests {
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
		if i == 0 {
			_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
				SpaceNumber: "A001",
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func TestAddVehicleHandler(t *testing.T) {
	var route = "/api/parking-spaces"
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
			description:   "Working request add parking space",
			expectedError: false,
			body:          "{\"space_number\":\"A001\"}",
			expectedCode:  201,
			expectedBody:  "{\"id\":1,\"space_number\":\"A001\"}",
		},
		{
			description:   "Bad request add parking space",
			expectedError: false,
			expectedCode:  400,
			body:          "",
			expectedBody:  "{\"error\":\"Cannot parse input\"}",
		},
		{
			description:   "Bad request add parking space",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"space_number\":\"vroooomm\"}",
			expectedBody:  "{\"error\":\"SpaceNumber is not valid (format: letter followed by a number (000 - 999))\"}",
		},
	}
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Post(route, AddParkingSpaceHandler)
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
