package ReservationController

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpace"
	"github.com/tot0p/CoursUT/internal/database/crud/parkingSpaceInformation"
	"github.com/tot0p/CoursUT/internal/database/crud/vehicle"
	"github.com/tot0p/CoursUT/internal/models"
	"io"
	"net/http"
	"testing"
	"time"
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

func TestGetReservationHandler(t *testing.T) {
	var route = "/api/reservations/:id"
	tests := []struct {
		description string

		// Test input
		id string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Valid reservation ID",
			id:            "1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"id\":1,\"vehicle_id\":1,\"parking_space_id\":1,\"arrival_time\":\"0001-01-01T00:00:00Z\",\"departure_time\":\"0001-01-01T00:00:00Z\",\"parking_duration\":3600000000000}",
		},
		{
			description:   "Invalid reservation ID",
			id:            "invalid",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Invalid ID\"}",
		},
		{
			description:   "Non-existent reservation ID",
			id:            "999",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Cannot get reservation\"}",
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

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpaceInformation.CreateParkingSpaceInformation(&models.ParkingSpaceInformation{
		ParkingSpaceID:  1,
		VehicleID:       1,
		ParkingDuration: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get(route, GetReservationHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			"/api/reservations/"+test.id,
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
func TestStartReservationHandler(t *testing.T) {
	var route = "/api/reservations/:id/start"
	tests := []struct {
		description   string
		id            string
		expectedError bool
		expectedCode  int
		expectedBody  string
		regexCheck    bool
	}{
		{
			description:   "Valid reservation ID",
			id:            "1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"id\":1,\"vehicle_id\":1,\"parking_space_id\":1,\"arrival_time\":\"" + time.Now().UTC().Format(time.RFC3339) + "\",\"departure_time\":\"0001-01-01T00:00:00Z\",\"parking_duration\":3600000000000}",
			regexCheck:    true,
		},
		{
			description:   "Invalid reservation ID",
			id:            "invalid",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Invalid reservation ID\"}",
		},
		{
			description:   "Non-existent reservation ID",
			id:            "999",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "{\"error\":\"Reservation not found\"}",
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

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpaceInformation.CreateParkingSpaceInformation(&models.ParkingSpaceInformation{
		ParkingSpaceID:  1,
		VehicleID:       1,
		ParkingDuration: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Post(route, StartReservationHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"POST",
			"/api/reservations/"+test.id+"/start",
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
		if test.regexCheck {
			// We can't predict the exact time, so we just check the format
			assert.Regexpf(t, `{"id":1,"vehicle_id":1,"parking_space_id":1,"arrival_time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,20}Z","departure_time":"0001-01-01T00:00:00Z","parking_duration":3600000000000}`, string(body), test.description)
		} else {
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}

func TestEndReservationHandler(t *testing.T) {
	var route = "/api/reservations/:id/end"
	tests := []struct {
		description   string
		id            string
		expectedError bool
		expectedCode  int
		expectedBody  string
		regexCheck    bool
	}{
		{
			description:   "Valid reservation ID",
			id:            "1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"id\":1,\"vehicle_id\":1,\"parking_space_id\":1,\"arrival_time\":\"0001-01-01T00:00:00Z\",\"departure_time\":\"" + time.Now().Format(time.RFC3339) + "\",\"parking_duration\":3600000000000}",
			regexCheck:    true,
		},
		{
			description:   "Invalid reservation ID",
			id:            "invalid",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Invalid reservation ID\"}",
		},
		{
			description:   "Non-existent reservation ID",
			id:            "999",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "{\"error\":\"Reservation not found\"}",
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

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpaceInformation.CreateParkingSpaceInformation(&models.ParkingSpaceInformation{
		ParkingSpaceID:  1,
		VehicleID:       1,
		ParkingDuration: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Post(route, EndReservationHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"POST",
			"/api/reservations/"+test.id+"/end",
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
		if test.regexCheck {
			// We can't predict the exact time, so we just check the format
			assert.Regexpf(t, `{"id":1,"vehicle_id":1,"parking_space_id":1,"arrival_time":"0001-01-01T00:00:00Z","departure_time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,20}Z","parking_duration":3600000000000}`, string(body), test.description)
		} else {
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}
func TestGetRemainingTimeHandler(t *testing.T) {
	var route = "/api/reservations/:id/remaining-time"
	tests := []struct {
		description   string
		id            string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Valid reservation ID",
			id:            "1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"remaining_time\":\"59m59s\"}",
		},
		{
			description:   "Invalid reservation ID",
			id:            "invalid",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Invalid reservation ID\"}",
		},
		{
			description:   "Non-existent reservation ID",
			id:            "999",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "{\"error\":\"Reservation not found\"}",
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

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpaceInformation.CreateParkingSpaceInformation(&models.ParkingSpaceInformation{
		ParkingSpaceID:  1,
		VehicleID:       1,
		ParkingDuration: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get(route, GetRemainingTimeHandler)
	app.Post("/api/reservations/:id/start", StartReservationHandler)

	// start the reservation
	req, _ := http.NewRequest(
		"POST",
		"/api/reservations/1/start",
		nil,
	)
	req.Header.Set("Content-Type", "application/json")
	_, _ = app.Test(req, -1)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			"/api/reservations/"+test.id+"/remaining-time",
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

func TestGetReservationQrCodeHandler(t *testing.T) {
	var route = "/api/reservations/:id/qrcode"
	tests := []struct {
		description   string
		id            string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "Valid reservation ID",
			id:            "1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "", // We can't predict the exact QR code content
		},
		{
			description:   "Invalid reservation ID",
			id:            "invalid",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":\"Invalid reservation ID\"}",
		},
		{
			description:   "Non-existent reservation ID",
			id:            "999",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "{\"error\":\"Reservation not found\"}",
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

	_, err = parkingSpace.CreateParkingSpace(&models.ParkingSpace{
		SpaceNumber: "A001",
		VehicleType: models.Car,
	})
	if err != nil {
		panic(err)
	}

	_, err = parkingSpaceInformation.CreateParkingSpaceInformation(&models.ParkingSpaceInformation{
		ParkingSpaceID:  1,
		VehicleID:       1,
		ParkingDuration: time.Hour,
	})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get(route, GetReservationQrCodeHandler)
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			"/api/reservations/"+test.id+"/qrcode",
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
		if test.expectedCode == 200 {
			assert.NotEmpty(t, body, test.description)
			assert.Nilf(t, err, test.description)
		} else {
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}
