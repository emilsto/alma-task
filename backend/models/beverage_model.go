package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO: Add validation

type Beverage struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Weight      float32            `json:"weight" validate:"required"`
	Price       float32            `json:"price" validate:"required"`
	RoastDegree int                `json:"roast_degree" validate:"required"` // Must be a number 1-5
	Type        string             `json:"type" validate:"required"`
}
