package repository

import (
	"context"
	"fmt"
	"letschat/infrastructure"
	"letschat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

func (c UserRepository) Create(user models.User) error {
	usersCollection := c.db.DB.Collection("users")
	_, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c UserRepository) Update(id string, user models.UpdateUser) error {
	usersCollection := c.db.DB.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": user}
	_, err := usersCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c UserRepository) Delete(id string) error {
	usersCollection := c.db.DB.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c UserRepository) FindOne(id string) (*models.User, error) {
	var user *models.User
	usersCollection := c.db.DB.Collection("users")

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (c UserRepository) FindAll() (*[]models.User, error) {
	var users []models.User
	usersCollection := c.db.DB.Collection("users")
	filter := bson.M{}
	cur, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &users, nil
}
