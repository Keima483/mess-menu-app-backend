package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/models"
	"github.com/keima483/mess-menu-app/services"
)

func GetMenuItems(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	items, err := services.GetMenuItem(id)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(items)
}

func AddMenuItem(c *fiber.Ctx) error {
	var itemDetails models.MenuItemModel
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&itemDetails); err != nil {
		return err
	}
	item, err := services.AddMenuItem(id, &itemDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(item)
}

func UpdateMenuItem(c *fiber.Ctx) error {
	var itemDetails models.MenuItemModel
	id, _ := strconv.Atoi(c.Params("id"))
	if err := c.BodyParser(&itemDetails); err != nil {
		return err
	}
	item, err := services.UpdateMenuItem(id, &itemDetails)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	return c.JSON(item)
}