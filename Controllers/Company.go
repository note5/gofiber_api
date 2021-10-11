package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"teleops/config"
	"teleops/models"
	"time"
)

//Create Company
func CreateCompany(c *fiber.Ctx) error {
	companyCollection := config.MI.DB.Collection("company")
	data := new(models.Company)
	validationError := data.Validate()

	if validationError != nil {
		fmt.Print(validationError)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Some fields failed in validation",
			"error":   validationError.Error(),
		})
	}

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
	time :=time.Now()
	data.CreatedAt = &time
	data.UpdatedAt = &time
    
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

//Get one company by id
func GetCompany(c *fiber.Ctx) error {
	companyCollection := config.MI.DB.Collection("company")
	// get parameter value
	paramID := c.Params("id")
	// convert parameterID to objectId
	id, err := primitive.ObjectIDFromHex(paramID)

	// if error while parsing paramID
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	//find company
	device := &models.Device{}
	query := bson.D{{Key: "_id", Value: id}}
	err = companyCollection.FindOne(c.Context(), query).Decode(device)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Company Not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": device,
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
	data := new(models.CompanyUpdate)
	err = c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	time :=time.Now()
	data.UpdatedAt = &time
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

//Delete company
func DeleteCompany(c *fiber.Ctx) error {

	collection := config.MI.DB.Collection("company")
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
	// find and delete todo
	query := bson.D{{Key: "_id", Value: id}}
	err = collection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Company Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete company",
			"error":   err,
		})
	}
	return c.SendStatus(fiber.StatusNoContent)

}