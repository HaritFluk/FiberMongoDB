package routes

import (
	"FiberMongoDB/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// All routes are here!!
	app.Post("/user", controllers.CreateUser) // Create a new user route

	app.Get("/user/:userId", controllers.GetAUser) // Get a user route

	app.Put("/user/:userId", controllers.EditAUser) // Edit a user route

	app.Delete("/user/:userId", controllers.DeleteAUser) // Delete a user route
}