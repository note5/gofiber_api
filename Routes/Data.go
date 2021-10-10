package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"teleops/services"
	"fmt"	
)

func DataRoute(route fiber.Router){
	Savedata(route ) //save  new device data
	GetAllData(route ) //save  new device data
	DeleteData(route ) //save  new device data
	GetDataByParams(route ) //save  new device data
}

//Save data to the databse

// Save data godoc
// @Summary Save device data
// @ID create-data
// @Description  Save device data
// @Accept  json
// @Produce  json
// @Tags Data End Points
// @param Data body models.Data true  "Data Details"
// @Success 200 {object} models.Data
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/data/create/ [post]
func Savedata(route fiber.Router){
	route.Post("/create/", services.CheckMiddleware, controllers.SaveData)
}

// GetAllData  godoc
// @Summary Get all Data
// @ID get-all-data
// @Description Get all Data
// @ID get-all-data
// @Accept  json
// @Produce  json
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Tags Data End Points
// @Success 200 {object} models.Data
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/data/ [get]
func GetAllData(route fiber.Router){
	route.Get("", services.CheckMiddleware, controllers.GetAllData)
}
// GetAllData  godoc
// @Summary Get all Data
// @ID filter-data-by-params
// @Description Filter data
// @ID filter-data-by-params
// @Accept  json
// @Produce  json
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Param device_addresses query string false "Comma seperated device addresses"
// @Tags Data End Points
// @Success 200 {object} models.Data
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/data/filter [get]
func GetDataByParams(route fiber.Router){
	route.Get("/filter", services.CheckMiddleware, controllers.GetDataByParams)
}

// Delete data by id godoc
// @Summary Delete record by id
// @ID delete-data-by-id
// @Description Delete data by id
// @Accept  json
// @Produce  json
// @Tags Data End Points
// @Param id path string true "Id" 
// @Success 200 {object} models.Data
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/data/{id} [delete]
func  DeleteData(route fiber.Router)  {
	fmt.Println("delete data ==============")
	route.Delete("/:id", services.CheckMiddleware, controllers.DeleteData)
}