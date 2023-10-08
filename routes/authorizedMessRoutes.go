package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/controllers"
	"github.com/keima483/mess-menu-app/middleware"
)

func AuthorizedMessRoutes(app *fiber.App) {
	group := app.Group("/api/v1/mess")
	group.Use(middleware.VerifyJWTToken)
	group.Post("", controllers.AddMess)
	group.Put("/:id", controllers.UpdateMess)
	group.Delete("/:id",controllers.DeleteMess)
}