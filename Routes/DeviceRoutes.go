package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"fmt"
)

func DeviceRoute(route fiber.Router) {
	route.Get("", controllers.GetDevices)
     route.Post("/create", checkMiddleware, controllers.CreateDevice)

}

//add custom middleware

func checkMiddleware(c *fiber.Ctx) error {
	fmt.Print("======== MIDDLEWARE RUNS======", string(c.Body()))
	// c.Next()
	return nil
}