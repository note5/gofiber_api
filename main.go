package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"teleops/config"
	"teleops/routes"
    // Import Fiber Swagger
    swagger "github.com/arsmn/fiber-swagger/v2"
	_ "teleops/docs"
		
)
// @title Teleops  API
// @version 2.0
// @description Teleops IOT server API
// @contact.name API Support
// @contact.email info@teleops.io
// @host localhost:3000
// @BasePath /

func main() {
	app := fiber.New()
	app.Use(logger.New())
	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
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
		// Add endpoint to serve swagger documentation
		app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			URL:         "/swagger/doc.json",
			DeepLinking: false,
		}))

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
	routes.DataRoute(api.Group("/data"))
}

type HTTPError struct {
	Status  string
	Message string
}