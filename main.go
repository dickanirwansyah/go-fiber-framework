package main

import (
	"github.com/dickanirwansyah/go-app/go-tech/database"
	"github.com/dickanirwansyah/go-app/go-tech/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectionDB()
	routes.Setup(app)
	app.Listen(":3000")
}
