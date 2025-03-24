package server

import (
	"github.com/gofiber/fiber/v2"
	apiHandler "github.com/tot0p/CoursUT/internal/server/controller/api"
	"github.com/tot0p/CoursUT/internal/server/controller/api/ParkingSpaceController"
	"github.com/tot0p/CoursUT/internal/server/controller/api/ReservationController"
	"github.com/tot0p/CoursUT/internal/server/controller/api/VehicleController"
)

func (serv *Server) handlers() {
	// api group
	api := serv.fiberApp.Group("/api")
	api.Get("/ping", apiHandler.PingHandler)
	api.Post("/vehicles", VehicleController.AddVehicleHandler)
	api.Get("/vehicles", VehicleController.GetVehiclesHandler)
	api.Delete("/vehicles/:id", VehicleController.DeleteVehicleHandler)
	api.Put("/vehicles/:id", VehicleController.UpdateVehicleHandler)

	api.Get("/parking-spaces", ParkingSpaceController.GetParkingSpacesHandler)
	api.Post("/parking-spaces", ParkingSpaceController.AddParkingSpaceHandler)
	api.Delete("/parking-spaces/:id", ParkingSpaceController.DeleteParkingSpacesHandler)
	api.Put("/parking-spaces/:id", ParkingSpaceController.UpdateParkingSpaceHandler)

	api.Get("/reservations/:id", ReservationController.GetReservationHandler)
	api.Post("/reservations", ReservationController.AddReservationHandler)
	api.Post("/reservations/:id/start", ReservationController.StartReservationHandler)
	api.Post("/reservations/:id/end", ReservationController.EndReservationHandler)
	api.Get("/reservations/:id/remaining-time", ReservationController.GetRemainingTimeHandler)
	api.Get("/reservations/:id/qrcode", ReservationController.GetReservationQrCodeHandler)
	api.Get("/reservations/:id/price", ReservationController.GetReservationPriceHandler)

	// Serve Swagger JSON
	serv.fiberApp.Get("/swagger/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger/swagger.json")
	})

	serv.fiberApp.Get("/swagger/", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger-ui/index.html")
	})
	serv.fiberApp.Get("/swagger/swagger-ui.js", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger-ui/swagger-ui.js")
	})
	serv.fiberApp.Get("/swagger/swagger-custom.css", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger-ui/swagger-custom.css")
	})
}
