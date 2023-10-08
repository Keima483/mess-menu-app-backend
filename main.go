package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/keima483/mess-menu-app/initializers"
	"github.com/keima483/mess-menu-app/routes"
)

func init() {
	config:= initializers.LoadConfig(".")
	log.Println(config)
	if err := initializers.InitializeDB(&config); err != nil {
		panic(err.Error())
	}
}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("refer to https://github.com/Keima483/fiber-ems for the endpoints")
	})
	routes.UnAuthorizedRoutes(app)
	routes.AuthorizedMessRoutes(app)
	routes.AuthorizedMenuItemRoutes(app)
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}
