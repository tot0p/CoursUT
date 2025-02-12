package server

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	serv := NewServer()
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	res, err := serv.fiberApp.Test(req)
	assert.Equalf(t, false, err != nil, "Expected status code 200, but got %v", res.StatusCode)
	assert.Equalf(t, 200, res.StatusCode, "Expected status code 200, but got %v", res.StatusCode)
	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, "Expected no error, but got %v", err)
	assert.Equalf(t, "{\"message\":\"pong\"}", string(body), "Expected body to be 'pong', but got %v", string(body))
}
