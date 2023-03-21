package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// the below lastmessageid will be message later on

type Room struct {
	RoomId        primitive.ObjectID `json:"room_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Createdby     string             `json:"created_by" bson:"created_by,omitempty"`
	LastMessageId string             `json:"last_message_id" bson:"last_message_id,omitempty"`
	Members       []string           `json:"members" bson:"members,omitempty"`
	ModifeidAt    time.Time          `json:"modifeid_at" bson:"modifeid_at,omitempty"`
	StartedAt     time.Time          `json:"started_at" bson:"started_at,omitempty"`
}

type RoomUpdate struct {
	Name          string    `json:"name" bson:"name,omitempty"`
	Createdby     string    `json:"created_by" bson:"created_by,omitempty"`
	LastMessageId string    `json:"last_message_id" bson:"last_message_id,omitempty"`
	Members       []string  `json:"members" bson:"members,omitempty"`
	ModifeidAt    time.Time `json:"modifeid_at" bson:"modifeid_at,omitempty"`
	StartedAt     time.Time `json:"started_at" bson:"started_at,omitempty"`
}
