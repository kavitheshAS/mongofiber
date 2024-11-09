package main

import (
	"mongofiber/configs"
	"mongofiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
