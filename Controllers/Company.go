package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"teleops/config"
	"teleops/models"
	"time"
)

//Create Company
func CreateCompany(c *fiber.Ctx) error {
	companyCollection := config.MI.DB.Collection("company")
	data := new(models.Company)
	err := c.BodyParser(&data)

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing JSON",
			"error":   err,
		})
	}
	data.ID = ""
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := companyCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"mssage":  "Error creating device",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"result":  result,
		"message": "Company created successfully",
	})

}

//Get All Companies

func GetCompanies(c *fiber.Ctx) error {
	companyCollection := config.MI.DB.Collection("company")
	// Query to filter
	query := bson.D{{}}
	cursor, err := companyCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not fetch companies",
			"error":   err.Error(),
		})
	}

	var companies []models.Company = make([]models.Company, 0)
	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &companies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Could not fetch companies",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": companies,
	})

}

//Update company details
func UpdateCompany(c *fiber.Ctx) error {

	companyCollection := config.MI.DB.Collection("company")
	// get param
	paramID := c.Params("id")
	// convert parameter to object id
	id, err := primitive.ObjectIDFromHex(paramID)
	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// var data Request
	data := new(models.Company)
	err = c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	data.UpdatedAt = time.Now()
	update := bson.M{
		"$set": data,
	}
	// fmt.Println(string(data)," updated values ",update)
	fmt.Printf("%+v\n", data)
	_, err = companyCollection.UpdateOne(c.Context(),  bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Could not update Company",
			"error":   err.Error(),
		})
	
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Company updated successfully",
	})

}