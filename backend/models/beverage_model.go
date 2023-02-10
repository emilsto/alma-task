package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

// TODO: Add validation

type Beverage struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Weight      float32            `json:"weight" validate:"required,gte=0,lte=10000"`
	Price       float32            `json:"price" validate:"required,gte=0,lte=10000"`    // Must be a number greater than or equal to 0
	RoastDegree int                `json:"roast_degree" validate:"required,gte=1,lte=5"` // Must be a number 1-5
	Type        string             `json:"type" validate:"required"`
}

func (b Beverage) Validate() error {
	return validator.New().Struct(b)
}
