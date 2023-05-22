package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Highlight struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	StartDate   string             `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate     string             `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Active      bool               `json:"active,omitempty" bson:"active,omitempty"`
}
