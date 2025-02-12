package server

import "github.com/gofiber/fiber/v2"

func (serv *Server) handlers() {
	serv.fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
