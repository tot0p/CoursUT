package VehicleController

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"io"
	"net/http"
	"testing"
)

func TestAddVehicleHandler(t *testing.T) {
	var route = "/api/vehicles"
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
			body:          "{\"plate\":\"AA-123-AA\",\"vehicleType\":1}",
			expectedCode:  201,
			expectedBody:  "{\"id\":1,\"plate\":\"AA-123-AA\",\"vehicleType\":1}",
		},
		{
			description:   "Bad request add vehicle",
			expectedError: false,
			expectedCode:  400,
			body:          "",
			expectedBody:  "{\"error\":\"Cannot parse input\"}",
		},
		{
			description:   "Bad request add vehicle without plate",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"\",\"vehicleType\":1}",
			expectedBody:  "{\"error\":\"Plate is required\"}",
		},
		{
			description:   "Bad request add vehicle without good type",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"AA-123-AA\",\"vehicleType\":100}",
			expectedBody:  "{\"error\":\"Type is required and must be different of Unknown\"}",
		},
		{
			description:   "Bad request add vehicle bad plate",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"VROOOOOOOM\",\"vehicleType\":1}",
			expectedBody:  "{\"error\":\"Plate is not valid\"}",
		},
	}
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Post(route, AddVehicleHandler)
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
			description:   "request Get vehicles",
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
			_, err = vehicle.CreateVehicle(&models.Vehicle{
				Plate:       "AA-123-AA",
				VehicleType: models.Car,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func TestUpdateVehicleHandler(t *testing.T) {
	var route = "/api/vehicles/:id"
	routeValue := "/api/vehicles/1"
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
			description:   "Working request Update vehicle",
			expectedError: false,
			body:          "{\"plate\":\"AA-123-AA\",\"vehicleType\":1}",
			expectedCode:  204,
			expectedBody:  "",
		},
		{
			description:   "Bad request update vehicle",
			expectedError: false,
			expectedCode:  400,
			body:          "",
			expectedBody:  "{\"error\":\"Cannot parse input\"}",
		},
		{
			description:   "Bad request update vehicle without plate",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"\",\"vehicleType\":1}",
			expectedBody:  "{\"error\":\"Plate is required\"}",
		},
		{
			description:   "Bad request update vehicle without good type",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"AA-123-AA\",\"vehicleType\":100}",
			expectedBody:  "{\"error\":\"Type is required and must be different of Unknown\"}",
		},
		{
			description:   "Bad request update vehicle bad plate",
			expectedError: false,
			expectedCode:  400,
			body:          "{\"plate\":\"VROOOOOOOM\",\"vehicleType\":1}",
			expectedBody:  "{\"error\":\"Plate is not valid\"}",
		},
	}
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	_, err = vehicle.CreateVehicle(&models.Vehicle{
		Plate:       "AA-123-AA",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}
	app.Put(route, UpdateVehicleHandler)
	for _, test := range tests {
		var bodyR io.Reader = nil
		if test.body != "" {
			bodyR = bytes.NewBuffer([]byte(test.body))
		}
		req, _ := http.NewRequest(
			"PUT",
			routeValue,
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
