package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"teleops/config"
	"teleops/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// config db
	config.ConnectDB()
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err = app.Listen(":3000")

	// handle error
	if err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {

	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	api := app.Group("/api")
	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})
	// connect device routes
	routes.DeviceRoute(api.Group("/devices"))
}
