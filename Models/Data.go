package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Device model

type Data struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	DeviceAddress      string             `json:"device_address,omitempty" validate:"required"`
	DeviceAlias        string             `json:"device_alias,omitempty"`
	LocationName       string             `json:"location_name,omitempty"`
	LocationId         string             `json:"location_id,omitempty"`
	GroupName          string             `json:"group_name,omitempty"`
	GroupId            string             `json:"group_id,omitempty"`
	OwnerId            string             `json:"owner_id,omitempty"`
	OwnerName          string             `json:"owner_name,omitempty"`
	PhysicalLocation   string             `json:"physical_location,omitempty"`
	CreatedById        string             `json:"created_by_id,omitempty"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at" swaggertype:"primitive,integer" swaggerignore:"true"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at" swaggertype:"primitive,integer" swaggerignore:"true"`
	SensorTypeName     string             `json:"sensor_type_name,omitempty"`
	SensorTypeId       string             `json:"sensor_type_id,omitempty"`
	SensorClassName    string             `json:"sensor_class_name,omitempty"`
	SensorClassId      string             `json:"sensor_class_id,omitempty"`
	IsDeployed         bool               `json:"is_deployed,omitempty"`
	HasGeolocationData bool               `json:"has_geolocation_data,omitempty"`
	Datetime           time.Time          `json:"datetime"  bson:"datetime"`

	Data []struct {
		Datetime       string `json:"datetime"`
		ProcessedValue string `json:"processed_value"`
		RawValue       string `json:"raw_value"`
		Geolocation    struct {
			Lat string `json:"lat"`
			Lon string `json:"lon"`
		} `json:"geolocation" bson:"geolocation"`
		Units      string `json:"units"`
		UnitString string `json:"unit_string"`
	}
}
