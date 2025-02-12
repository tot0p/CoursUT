package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	fiberApp *fiber.App
}

func NewServer() *Server {
	serv := &Server{
		fiber.New(),
	}
	serv.fiberApp.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Local",
	}))
	return serv
}

func (serv *Server) Run() {
	serv.handlers()
	if err := serv.fiberApp.Listen(":8080"); err != nil {
		panic(err)
	}
}
