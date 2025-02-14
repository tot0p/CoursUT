package VehicleController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestPingHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/api/ping", PingHandler)
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	res, err := app.Test(req)
	assert.Equalf(t, false, err != nil, "Expected status code 200, but got %v", res.StatusCode)
	assert.Equalf(t, 200, res.StatusCode, "Expected status code 200, but got %v", res.StatusCode)
	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, "Expected no error, but got %v", err)
	assert.Equalf(t, "{\"message\":\"pong\"}", string(body), "Expected body to be 'pong', but got %v", string(body))
}
