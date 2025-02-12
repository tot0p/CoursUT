package server

import (
	apiHandler "github.com/tot0p/CoursUT/internal/server/controller/api"
)

func (serv *Server) handlers() {

	// api group
	api := serv.fiberApp.Group("/api")
	api.Get("/ping", apiHandler.PingHandler)
	api.Post("/vehicles", apiHandler.AddVehicleHandler)
	api.Get("/vehicles", apiHandler.GetVehiclesHandler)
}
