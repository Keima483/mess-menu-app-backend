package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/controllers"
)

func UnAuthorizedRoutes(app *fiber.App) {
	group := app.Group("/api/v1")
	group.Post("/mess/login", controllers.Login)
	group.Get("/mess/", controllers.GetMess)
	group.Get("/mess/:id", controllers.GetMessById)
	group.Get("/item/:id", controllers.GetMenuItems)
}