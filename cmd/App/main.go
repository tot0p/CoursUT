package main

import (
	"github.com/tot0p/CoursUT/internal/database"
	"github.com/tot0p/CoursUT/internal/server"
)

// @title CoursUT API
// @version 1.0
// @description This is the API documentation for the CoursUT parking management system.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your.email@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	err := database.InitDatabase()
	if err != nil {
		panic(err)
	}
	server.NewServer().Run()
}
