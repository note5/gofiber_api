package models

import (
	"time"
)

//Device model

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


