package repository

import (
	"context"
	"fmt"
	"letschat/infrastructure"
	"letschat/models"

	"go.mongodb.org/mongo-driver/bson"
)

type RoomRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewRoomRepository(db infrastructure.Database, logger infrastructure.Logger) RoomRepository {
	return RoomRepository{
		db:     db,
		logger: logger,
	}
}

func (c RoomRepository) Create(room models.Room) error {
	roomsCollection := c.db.DB.Collection("rooms")
	_, err := roomsCollection.InsertOne(context.TODO(), room)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomRepository) Update(id string, room models.Room) error {
	roomsCollection := c.db.DB.Collection("rooms")
	filter := bson.M{"display_name": "kapil"}
	update := bson.M{"$set": room}
	_, err := roomsCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomRepository) Delete(id string) error {
	roomsCollection := c.db.DB.Collection("rooms")
	filter := bson.M{"room_id": id}
	_, err := roomsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomRepository) FindOne(id string) (*models.Room, error) {
	var room *models.Room
	roomsCollection := c.db.DB.Collection("rooms")
	filter := bson.M{"room_id": id}
	err := roomsCollection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return room, nil
}
