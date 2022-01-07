package controller

import (
	"log"

	"github.com/dickanirwansyah/go-app/go-tech/database"
	"github.com/dickanirwansyah/go-app/go-tech/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password yang anda masukan tidak match dengan password sebelumnya",
			"status":  400,
			"path":    c.BaseURL() + c.Path(),
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	database.DB.Create(&user)
	log.Print("response data : ", user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "username & password tidak ditemukan",
			"status":  404,
			"path":    c.BaseURL() + c.Path(),
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password yang anda masukan salah",
			"status":  400,
			"path":    c.BaseURL() + c.Path(),
		})
	}

	return c.JSON(user)
}
