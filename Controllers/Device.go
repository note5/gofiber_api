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

//Create Device
func CreateDevice(c *fiber.Ctx) error {
	deviceCollection := config.MI.DB.Collection("devices")

	data := new(models.Device)

	err := c.BodyParser(&data)

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error parsing JSON",
			"error":   err,
		})
	}
	fmt.Print("=========== ID===========", data.ID)
    data.ID = ""
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := deviceCollection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"mssage":  "Error creating device",
			"error":   err,
		})
	}

	//get the inserte data

	device := &models.Device{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	deviceCollection.FindOne(c.Context(), query).Decode(device)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"result":  device,
	})

}

// Get all devices

func GetDevices(c *fiber.Ctx) error {
	deviceCollection := config.MI.DB.Collection("devices")

	// Query to filter
	query := bson.D{{}}

	cursor, err := deviceCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err,
		})
	}

	var devices []models.Device = make([]models.Device, 0)

	// iterate the cursor anddecode each item into a Todo
	err = cursor.All(c.Context(), &devices)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": devices,
	})
}

//Get One Device

func GetDevice(c *fiber.Ctx) error {
	deviceCollection := config.MI.DB.Collection("devices")
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

	//find device
	device := &models.Device{}
	query := bson.D{{Key: "_id", Value: id}}
	err = deviceCollection.FindOne(c.Context(), query).Decode(device)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Device Not found",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"results": device,
	})

}

//Delete device

func DeleteDevice(c *fiber.Ctx) error {

	deviceCollection := config.MI.DB.Collection("devices")
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
	err = deviceCollection.FindOneAndDelete(c.Context(), query).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Device Not found",
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

//Update device
func UpdateDevice(c *fiber.Ctx) error {

	deviceCollection := config.MI.DB.Collection("devices")
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
	data := new(models.Device)
	err = c.BodyParser(&data)
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
	_, err = deviceCollection.UpdateOne(c.Context(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Could not update Device",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Device updated successfully",
	})

}
