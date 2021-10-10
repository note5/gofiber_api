package models

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

//Device model

type Device struct {
	ID               string    `json:"id,omitempty" bson:"_id,omitempty" `
	DeviceAddress    string    `json:"device_address,omitempty" bson:"device_address" validate:"required"`
	DeviceAlias      string    `json:"device_alias,omitempty" bson:"device_alias"`
	LocationName     string    `json:"location_name,omitempty" bson:"location_name"`
	LocationId       string    `json:"location_id,omitempty" bson:"location_id"`
	GroupName        string    `json:"group_name,omitempty" bson:"group_name"`
	GroupId          string    `json:"group_id,omitempty" bson:"group_id"`
	OwnerId          string    `json:"owner_id,omitempty" bson:"owner_id"`
	OwnerName        string    `json:"owner_name,omitempty" bson:"owner_name"`
	PhysicalLocation string    `json:"physical_location,omitempty" bson:"physical_location"`
	CreatedByEmail   string    `json:"created_by_email,omitempty" bson:"created_by_email"`
	CreatedById      string    `json:"created_by_id,omitempty" bson:"created_by_id"`
	CreatedAt        time.Time `json:"created_at" bson:"created_at"  swaggerignore:"true" `
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at"  swaggerignore:"true"`
	SensorTypeName   string    `json:"sensor_type_name,omitempty" bson:"sensor_type_name"`
	SensorTypeId     string    `json:"sensor_type_id,omitempty" bson:"sensor_type_id"`
	SensorClassName  string    `json:"sensor_class_name,omitempty" bson:"sensor_class_name"`
	SensorClassId    string    `json:"sensor_class_id,omitempty" bson:"sensor_class_id"`

	IsDeployed         bool `json:"is_deployed,omitempty" bson:"is_deployed"`
	HasGeolocationData bool `json:"has_geolocation_data,omitempty" bson:"has_geolocation_data"`

	ClientDetails []struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		CanGetSms bool   `json:"can_get_sms" bson:"can_get_sms"`
	} `json:"client_details" bson:"client_details"`
	Geolocation struct {
		Lat string `json:"lat,omitempty"`
		Lon string `json:"lon,omitempty"`
	}  `json:"geolocation"bson:"geolocation"`
}



