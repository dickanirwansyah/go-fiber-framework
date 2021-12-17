package controller

import (
	"log"

	"github.com/dickanirwansyah/go-app/go-tech/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Muhammad",
		LastName:  "Dicka nirwansyah",
		Email:     "dickanirwansyah@gmail.com",
		Password:  "P@sw0rdbb1",
	}
	log.Print("response data : ", user)
	return c.JSON(user)
}
