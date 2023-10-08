package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/initializers"
	"github.com/keima483/mess-menu-app/middleware"
	"github.com/keima483/mess-menu-app/models"
	"github.com/keima483/mess-menu-app/services"
)

func Login(c *fiber.Ctx) error {
	config := initializers.LoadConfig(".")
	loginDetails := new(models.LoginModel)
	if err := c.BodyParser(loginDetails); err != nil {
		return err
	}
	if config.AdminEmail != loginDetails.Email || config.AdminPassword != loginDetails.Password {
		return fiber.NewError(fiber.StatusBadRequest, "Wrong Email or password")
	}
	token, exp, err := middleware.CreateJWTToken()
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token, "exp": exp})
}

func GetMess(c *fiber.Ctx) error {
	return c.JSON(services.GetMess())
} 

func GetMessById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	mess, err := services.GetMessById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(mess)
}

func AddMess(c *fiber.Ctx) error {
	var messDetails models.MessModel
	if err := c.BodyParser(&messDetails); err != nil {
		return err
	}
	mess, err := services.AddMess(&messDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(mess)
}

func UpdateMess(c *fiber.Ctx) error {
	var messDetails models.MessModel
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&messDetails); err != nil {
		return err
	}
	mess, err := services.UpdateMess(id, messDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(mess)
}

func DeleteMess(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := services.DeleteMess(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.SendString("Successfully deleted")
}

