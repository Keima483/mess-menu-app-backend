package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/controllers"
	"github.com/keima483/mess-menu-app/middleware"
)


func AuthorizedMenuItemRoutes(app *fiber.App) {
	group := app.Group("/api/v1/item")
	group.Use(middleware.VerifyJWTToken)
	group.Post("/:id", controllers.AddMenuItem)
	group.Put("/:id", controllers.UpdateMenuItem)
}