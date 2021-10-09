package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"teleops/services"	
)
//main entry point for route
func DeviceRoute(route fiber.Router) {
	GetAllDevices(route ) //get add devices
	CreateDevice(route ) //create a new device
	GetDevice(route) //get one device
	DeleteDevice(route) //delete one device
	UpdateDevice(route) //update one device
	
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

// Get One Device godoc
// @Summary Get device by id
// @ID get-one-device
// @Description Get device by id
// @Accept  json
// @Produce  json
// @Tags Devices End Points
// @Param id path string true "Device Id" 
// @Success 200 {object} models.Device
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/devices/{id} [get]
func  GetDevice(route fiber.Router)  {
	route.Get("/:id", services.CheckMiddleware, controllers.GetDevice)
}
// Delete One Device godoc
// @Summary Delete device by id
// @ID delete-one-device
// @Description Delete device by id
// @Accept  json
// @Produce  json
// @Tags Devices End Points
// @Param id path string true "Device Id" 
// @Success 200 {object} models.Device
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/devices/{id} [delete]
func  DeleteDevice(route fiber.Router)  {
	route.Delete("/:id", services.CheckMiddleware, controllers.DeleteDevice)
}
// Update One Device godoc
// @Summary Update device by id
// @ID update-one-device
// @Description Update device by id
// @Accept  json
// @Produce  json
// @Tags Devices End Points
// @Param id path string true "Device Id" 
// @param device body models.Device true  "Device Details"
// @Success 200 {object} models.Device
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/devices/{id} [patch]
func  UpdateDevice(route fiber.Router)  {
	route.Patch("/:id", services.CheckMiddleware, controllers.UpdateDevice)
}


