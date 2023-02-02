package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Beverage struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	WeightGrams int                `json:"weight_grams " validate:"required"`
	Price       float32            `json:"price" validate:"required"`
	RoastDegree int                `json:"roast_degree" validate:"required"` // Must be a nubmer 1-5 TODO: Add validation
	Type        string             `json:"type" validate:"required"`
}
