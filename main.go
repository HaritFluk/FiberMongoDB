package main

import (

	"FiberMongoDB/configs"
	"FiberMongoDB/routes"

	"github.com/gofiber/fiber/v2"
	)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	// run routes
	routes.UserRoute(app)

	app.Listen(":6000")
}