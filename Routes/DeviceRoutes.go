package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"fmt"
)

func DeviceRoute(route fiber.Router) {
	route.Get("", controllers.GetDevices)
     route.Post("/create", checkMiddle, controllers.CreateDevice)

}

func checkMiddle(){
	fmt.Print("======== MIDDLEWARE RUNS======")
}