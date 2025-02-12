package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	fiberApp *fiber.App
}

func NewServer() *Server {
	return &Server{
		fiber.New(),
	}
}

func (serv *Server) Run() {
	serv.handlers()
	if err := serv.fiberApp.Listen(":8080"); err != nil {
		panic(err)
	}
}
