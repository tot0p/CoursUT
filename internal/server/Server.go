package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
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
	serv.fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	serv.handlers()
	return serv
}

func (serv *Server) Run() {
	log.Fatal(serv.fiberApp.Listen(":8080"))
}
