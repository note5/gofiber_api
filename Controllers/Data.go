package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"teleops/config"
	"teleops/models"
	"time"
	// 	"go.mongodb.org/mongo-driver/bson/primitive"
	// 	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Save Data
func SaveData(c *fiber.Ctx) error {
	dataCollection := config.MI.DB.Collection("data") //get the collection of the data
	data := new(models.Data)                          //instance of the Data struct
	err := c.BodyParser(&data)                        //unmarshal the json to do struct
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing JSON data",
			"error":   err,
		})
	}
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := dataCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"mssage":  "Error Saving data",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Data inserted successfully",
		"result":  result,
	})

}

//Get all data

func GetAllData(c *fiber.Ctx) error {
	dataCollection := config.MI.DB.Collection("data")                 //get the collection of the data
	startDate, err := time.Parse(time.RFC3339, c.Query("start_date")) //get start_date
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing Start Date",
			"error":   err,
		})
	}
	endDate, err := time.Parse(time.RFC3339, c.Query("end_date"))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing End Date",
			"error":   err,
		})
	}
	//get end_date
	// fmt.Print(startDate, "===================== ", endDate)
	query := bson.M{
		"datetime": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	cursor, err := dataCollection.Find(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}
	var data []models.Data = make([]models.Data, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error  getting data",
			"error":   err.Error(),
		})
	}

	fmt.Println("Length Data , ", len(data) > 0)

	if len(data) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"results": data,
		})
	}
	//query
	options := options.Find()
	query = bson.M{}
	options.SetSort(bson.D{{"datetime", -1}})
	options.SetLimit(4)
	cursor, err = dataCollection.Find(c.Context(), query,options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}
	err = cursor.All(c.Context(), &data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error  getting data",
			"error":   err.Error(),
		})
	}

	fmt.Println("Length Data second , ", len(data) > 0)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": data,
		"length": len(data),
	})

}
