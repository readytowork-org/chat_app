package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId      primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	DisplayName string             `json:"display_name" bson:"display_name,omitempty"`
	Status      bool               `json:"status" bson:"status,omitempty"`
	Address     string             `json:"address" bson:"address,omitempty"`
	PhoneNumber string             `json:"phone" bson:"phone,omitempty"`
	PhotoUrl    string             `json:"photo" bson:"photo,omitempty"`
	Rooms       []string           `json:"rooms" bson:"rooms,omitempty"`
}
