package controllers

import (
	"fmt"
	"strings"
	"teleops/config"
	"teleops/models"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
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
	//check if we have start date if not just get limited entries
	if  c.Query("start_date") =="" {
		return GetLimitedData(c) 
	}
	//get the collection of the data
	dataCollection := config.MI.DB.Collection("data")
	//get start_date                 
	 startDate, err := time.Parse(time.RFC3339, c.Query("start_date")) 
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

	//if we have data
	if len(data) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"results": data,
		})
	}
	return GetLimitedData(c)

}

//Get data by more params like datetime range and id or device_alias
func GetDataByParams(c *fiber.Ctx) error {
	dataCollection := config.MI.DB.Collection("data")
	//
	options := options.Find()
	options.SetSort(bson.D{{"datetime", -1}})
	startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(c.Query("start_date"))) //get start_date
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing Start Date",
			"error":   err,
		})
	}
	endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(c.Query("end_date")))
	//split the parameters by , to get a slice
	params := strings.Split(strings.TrimSpace(c.Query("params")), ",")
	param := strings.TrimSpace(c.Query("param"))
	owner_id := strings.TrimSpace(c.Query("owner_id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing End Date",
			"error":   err,
		})
	}
	query := bson.M{
        "datetime": bson.M{
            "$gt": startDate,
            "$lt": endDate,
        },
		"owner_id": owner_id,
        param:bson.M{"$in": params},
    }
	fmt.Println(" Query ", params)
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": data,
	})
}

//Delete data by id
func DeleteData(c *fiber.Ctx) error {

	Collection := config.MI.DB.Collection("data")
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
	err = Collection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Data Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete todo",
			"error":   err,
		})
	}
	return c.SendStatus(fiber.StatusNoContent)

}

//get data and limit when query condition fails
func GetLimitedData(c *fiber.Ctx) error {
	dataCollection := config.MI.DB.Collection("data")
	options := options.Find()
	query := bson.M{}
	options.SetSort(bson.D{{"datetime", -1}})
	options.SetLimit(10)
	cursor, err := dataCollection.Find(c.Context(), query, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}
	var data []models.Data = make([]models.Data, 0)
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
		"length":  len(data),
	})

}
