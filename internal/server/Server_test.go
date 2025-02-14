package server

import (
	"net/http"
	"testing"
)

func TestNewServer(t *testing.T) {
	var s *Server
	s = NewServer()
	if s == nil {
		t.Errorf("NewServer() = nil, want not nil")
		return
	}
	if s.fiberApp == nil {
		t.Errorf("NewServer() = nil, want not nil")
	}
}

func TestServer_Run(t *testing.T) {
	s := NewServer()
	go func() {
		s.Run()
	}()

	// make request to server at /api/ping
	// check if response is 200
	req, err := http.NewRequest("GET", "http://localhost:8080/api/ping", nil)
	if err != nil {
		t.Errorf("Error creating request")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
