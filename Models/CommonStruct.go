package models

import (
	"time"
)

//Common model for tracking who makes chnages and the time changes were made

type CommonItems struct {
	
	CompanyId      string    `json:"company_id,omitempty" bson:"company_id"`
	CompanyName    string    `json:"company_name,omitempty" bson:"company_name"`
	CreatedByEmail string    `json:"created_by_email,omitempty" bson:"created_by_email"`
	CreatedById    string    `json:"created_by_id,omitempty" bson:"created_by_id"`
	CreatedAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
