package main

import (
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/server"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	server.NewServer().Run()
}
