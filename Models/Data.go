package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Device model

type Data struct {
	ID               primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
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
	CreatedAt        time.Time `json:"created_at" bson:"created_at" swaggertype:"primitive,integer" swaggerignore:"true " `
	UpdatedAt        time.Time `json:"updated_at" bson:"updated_at" swaggertype:"primitive,integer" swaggerignore:"true"`
	SensorTypeName   string    `json:"sensor_type_name,omitempty" bson:"sensor_type_name"`
	SensorTypeId     string    `json:"sensor_type_id,omitempty" bson:"sensor_type_id"`
	SensorClassName  string    `json:"sensor_class_name,omitempty" bson:"sensor_class_name"`
	SensorClassId    string    `json:"sensor_class_id,omitempty" bson:"sensor_class_id"`
	IsDeployed         bool `json:"is_deployed,omitempty" bson:"is_deployed"`
	HasGeolocationData bool `json:"has_geolocation_data,omitempty" bson:"has_geolocation_data"`
	Datetime           time.Time          `json:"datetime"  bson:"datetime"`

	Data []struct {
		Parameter string `json:"parameter" bson:"parameter"`
		Datetime       string `json:"datetime" bson:"datetime"`
		ProcessedValue string `json:"processed_value" bson:"processed_value"`
		RawValue       string `json:"raw_value" bson:"raw_value"`
		Geolocation    struct {
			Lat string `json:"lat"`
			Lon string `json:"lon"`
		} `json:"geolocation" bson:"geolocation"`
		Units      string `json:"units" bson:"units"`
		UnitString string `json:"unit_string" bson:"unit_string"`
	}
}
