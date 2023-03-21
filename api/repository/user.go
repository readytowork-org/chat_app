package repository

import (
	"context"
	"errors"
	"fmt"
	"letschat/infrastructure"
	"letschat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c UserRepository) Create(user models.CreateUser) error {
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
	println(filter)
	update := bson.M{"$set": user}
	result, err := usersCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("Cannot Update: No Document Found With The Id")
	}
	fmt.Println(result.ModifiedCount, "documents updated")
	return nil
}

func (c UserRepository) Delete(id string) error {
	usersCollection := c.db.DB.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("Cannot Delete: No Document Found With The id")
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
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Cannot Find User: No Document Found With The id")
		}
		return nil, err
	}
	return user, nil
}

func (c UserRepository) CheckUserWithPhone(phone string) (*models.CreateUser, bool, error) {
	var user *models.CreateUser
	usersCollection := c.db.DB.Collection("users")
	filter := bson.M{"phone": phone}
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching documents found")
			return nil, false, nil
		} else {
			fmt.Println("Error:", err)
			return nil, true, err
		}
	}
	return user, true, nil
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
