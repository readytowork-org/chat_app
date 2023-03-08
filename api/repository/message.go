package repository

import (
	"context"
	"fmt"
	"letschat/infrastructure"
	"letschat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewMessageRepository(db infrastructure.Database, logger infrastructure.Logger) MessageRepository {
	return MessageRepository{
		db:     db,
		logger: logger,
	}
}

func (c MessageRepository) Create(message models.MessageM) error {
	usersCollection := c.db.DB.Collection("messages")
	_, err := usersCollection.InsertOne(context.TODO(), message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c MessageRepository) Delete(id string) error {
	usersCollection := c.db.DB.Collection("messages")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c MessageRepository) FindAll(roomId string) (*[]models.MessageM, error) {
	var messages []models.MessageM
	usersCollection := c.db.DB.Collection("messages")
	filter := bson.M{"room_id": roomId}
	cur, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var message models.MessageM
		err := cur.Decode(&message)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &messages, nil
}
