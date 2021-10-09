package controllers

import (
	"fmt"
	"encoding/json"
	"time"
    "github.com/gofiber/fiber/v2"
    "teleops/models"
    "teleops/config"
    "go.mongodb.org/mongo-driver/bson" // new
    // "go.mongodb.org/mongo-driver/bson/primitive" // nw
    // "go.mongodb.org/mongo-driver/mongo" // new
)

type Device struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty"`
	DeviceAddress    string    `json:"device_address,omitempty"`
	DeviceAlias      string    `json:"device_alias,omitempty"`
	LocationName     string    `json:"location_name,omitempty"`
	LocationId       string    `json:"location_id,omitempty"`
	GroupName        string    `json:"group_name,omitempty"`
	GroupId          string    `json:"group_id,omitempty"`
	OwnerId          string    `json:"owner_id,omitempty"`
	OwnerName        string    `json:"owner_name,omitempty"`
	PhysicalLocation string    `json:"physical_location,omitempty"`
	CreatedByEmail   string    `json:"created_by_email,omitempty"`
	CreatedById      string    `json:"created_by_id,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	SensorTypeName   string    `json:"sensor_type_name,omitempty"`
	SensorTypeId     string    `json:"sensor_type_id,omitempty"`
	SensorClassName  string    `json:"sensor_class_name,omitempty"`
	SensorClassId    string    `json:"sensor_class_id,omitempty"`

	IsDeployed         bool `Json:"is_deployed,omitempty,omitempty"`
	HasGeolocationData bool `Json:"has_geolocation_data,omitempty"`

	ClientDetails []struct {
		Name      string `Json:"name"`
		Email     string `Json:"email"`
		Phone     string `Json:"phone"`
		CanGetSms bool   `Json:"can_get_sms"`
	} `Json:"client_details"`
	Geolocation struct{
		Lat string `Json:"lat,omitempty"`
		Lon string `Json:"lon,omitempty"`
	} `Json:"geolocation,omitempty"`
}

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
            "error":   err.Error(),
       })
    }

   var devices []models.Device = make([]models.Device, 0)

    // iterate the cursor anddecode each item into a Todo
    err = cursor.All(c.Context(), &devices)
    if err != nil {
       return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
           "success": false,
            "message": "Soething went wrong",
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
	var d Device
	json.Unmarshal(c.Body(),&d)
	
	fmt.Printf("%s\n", data.ClientDetails)
	
	err :=c.BodyParser(&d)
	
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