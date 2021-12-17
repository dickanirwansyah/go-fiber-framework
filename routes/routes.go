package routes

import (
	"github.com/dickanirwansyah/go-app/go-tech/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/register", controller.Register)
}
