package models

type Company struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty" `
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CommonItems        //commmon items struct

}
