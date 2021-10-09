package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
)

func DeviceRoute(route fiber.Router) {
	route.Get("", controllers.GetDevices)
route.Post("/create", controllers.CreateDevice)

}
