package models

import (
	"time"
	"github.com/go-playground/validator/v10"
)

type Company struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	IsActive    *bool   `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CompanyId      string    `json:"company_id,omitempty" bson:"company_id"`
	CompanyName    string    `json:"company_name,omitempty" bson:"company_name"`
	CreatedByEmail string    `json:"created_by_email,omitempty" bson:"created_by_email"`
	CreatedById    string    `json:"created_by_id,omitempty" bson:"created_by_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`

}

type CompanyUpdate struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty" `
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CompanyId      string    `json:"company_id,omitempty" bson:"company_id"`
	CompanyName    string    `json:"company_name,omitempty" bson:"company_name"`
	UpdatedByEmail string    `json:"created_by_email,omitempty" bson:"created_by_email"`
	UpdatedById    string    `json:"created_by_id,omitempty" bson:"created_by_id"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (c *Company) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	return err
}