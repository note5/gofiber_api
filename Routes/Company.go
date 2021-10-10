package routes

import (
	"github.com/gofiber/fiber/v2"
	"teleops/controllers"
	"teleops/services"
)

//main entry point for route
func CompanyRoute(route fiber.Router) {
	CreateCompany(route) //create a new company
	GetCompanies(route) //get all companies
	UpdateCompany(route) //update company


}

// Create Company godoc
// @Summary Create a Company
// @ID create-company
// @Description  Create a device
// @Accept  json
// @Produce  json
// @Tags Company End Points
// @param device body models.Company true  "Company Details"
// @Success 200 {object} models.Company
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/companies/create [post]
func CreateCompany(route fiber.Router) {
	route.Post("/create", services.CheckMiddleware, controllers.CreateCompany)
}

// Get Companies godoc
// @Summary Get all Companies
// @ID get-all-companies
// @Description Get all Companies
// @ID get-all-companies
// @Accept  json
// @Produce  json
// @Tags Company End Points
// @Success 200 {object} models.Company
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/companies/ [get]
func GetCompanies(route fiber.Router){
	route.Get("", controllers.GetCompanies)
}


// Update One Company godoc
// @Summary Update company by id
// @ID update-one-company
// @Description Update company by id
// @Accept  json
// @Produce  json
// @Tags Company End Points
// @Param id path string true "Company Id" 
// @param Company_details body models.Company false  "Company Details"
// @Success 200 {object} models.Company
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /api/companies/{id} [patch]
func  UpdateCompany(route fiber.Router)  {
	route.Patch("/:id", services.CheckMiddleware, controllers.UpdateCompany)
}
