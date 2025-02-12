package server

import (
	apiHandler "github.com/tot0p/CoursUT/internal/server/controller/api"
)

func (serv *Server) handlers() {

	// api group
	api := serv.fiberApp.Group("/api")
	api.Get("/ping", apiHandler.GetPingHandler())
	api.Post("/vehicles", apiHandler.GetAddVehicleHandler())
}
