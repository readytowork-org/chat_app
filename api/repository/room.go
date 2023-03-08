package repository

import (
	"context"
	"fmt"
	"letschat/infrastructure"
	"letschat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c RoomRepository) Update(id string, room models.RoomUpdate) error {
	roomsCollection := c.db.DB.Collection("rooms")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

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
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
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
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	as := fmt.Sprint(objID)
	println(as)
	err := roomsCollection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return room, nil
}
