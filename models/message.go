package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Message  string `json:"message"`
	Type     string `json:"type"`
	ClientID string `json:"clientId"`
}

type MessageM struct {
MessageId  primitive.ObjectID `json:"message_id" bson:"_id,omitempty"`
Message    string             `json:"message"`
Created_At time.Time          `json:"created_at"`
Created_By string             `json:"created_by"`
Type       string             `json:"type"`
Room_Id    string             `json:"room_id"`
}
