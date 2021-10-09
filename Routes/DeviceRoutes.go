package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"teleops/services"
	
)
//main entry point for route
func DeviceRoute(route fiber.Router) {
	GetAllDevices(route )
	CreateDevice(route )
	
	}
	
// GetDevices godoc
// @Summary Get all Devices
// @ID get-all-devices
// @Description Get all Devices
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Devices End Points
// @Success 200 {object} models.Device
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/devices/ [get]
func GetAllDevices(route fiber.Router){
	route.Get("", controllers.GetDevices)
}
// Create Device godoc
// @Summary Create a device
// @ID create-device
// @Description  Create a device
// @Accept  json
// @Produce  json
// @Tags Devices End Points
// @param device body models.Device true  "Device Details"
// @Success 200 {object} models.Device
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/devices/create [post]
func CreateDevice(route fiber.Router){
	route.Post("/create", services.CheckMiddleware, controllers.CreateDevice)
}



