package controllers

import (
	"fmt"
	
	"time"
    "github.com/gofiber/fiber/v2"
    "teleops/models"
    "teleops/config"
    "go.mongodb.org/mongo-driver/bson" // new
    // "go.mongodb.org/mongo-driver/bson/primitive" // nw
    // "go.mongodb.org/mongo-driver/mongo" // new
)



// get all devices
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
        "results":devices,
    })
}

func CreateDevice(c *fiber.Ctx) error{
	deviceCollection := config.MI.DB.Collection("devices")
    
	data := new(models.Device)
	
	err :=c.BodyParser(&data)
	
	if err !=nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"Error parsing JSON",
			"error":err,
		})
	}
	
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := deviceCollection.InsertOne(c.Context(),data)
	if err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"mssage":"Error creating device",
			"error":err,
		})
	}
	

	//get the inserte data 

	device := &models.Device{}
	query := bson.D{{Key:"_id",Value :result.InsertedID}}

	deviceCollection.FindOne(c.Context(),query).Decode(device)
 
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	"success":true,
		"result":device,
	})

	}