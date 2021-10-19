package routes

import (
	"shortner/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes (app *fiber.App){
	app.Post("/shorten", controllers.Shortner)
	app.Get("/:code", controllers.Redirect)
}