package main

import (
	"fmt"
	"log"

	"shortner/config"
	"shortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)


func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main(){

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.SetupRoutes(app)

	PORT := config.GetEnv("PORT", "8000") 

	fmt.Println("Application started...")
	fmt.Println(PORT)
	app.Listen(fmt.Sprintf(":%s", PORT))
}