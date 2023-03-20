package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type BaseUser struct {
	DisplayName string `json:"display_name" bson:"display_name,omitempty"`
	Email       string `json:"email" bson:"email,omitempty"`
	PhoneNumber string `json:"phone" bson:"phone,omitempty" validate:"required"`
}

type CreateUser struct {
	UserId          primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	DisplayName     string             `json:"display_name" bson:"display_name,omitempty"`
	Email           string             `json:"email" bson:"email,omitempty"`
	PhoneNumber     string             `json:"phone" bson:"phone,omitempty" validate:"required,phone"`
	Password        string             `json:"password" bson:"password" validate:"required,password"`
	ConfirmPassword string             `json:"confirm_password" validate:"required,confirm_password"`
}

type User struct {
	UserId primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	BaseUser
	Status   bool     `json:"status" bson:"status,omitempty"`
	Address  string   `json:"address" bson:"address,omitempty"`
	PhotoUrl string   `json:"photo" bson:"photo,omitempty"`
	Rooms    []string `json:"rooms" bson:"rooms,omitempty"`
}
type UpdateUser struct {
	Email       string   `json:"email" bson:"email,omitempty"`
	DisplayName string   `json:"display_name" bson:"display_name,omitempty"`
	Status      bool     `json:"status" bson:"status,omitempty"`
	Address     string   `json:"address" bson:"address,omitempty"`
	PhoneNumber string   `json:"phone" bson:"phone,omitempty"`
	PhotoUrl    string   `json:"photo" bson:"photo,omitempty"`
	Rooms       []string `json:"rooms" bson:"rooms,omitempty"`
}

func (u CreateUser) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":      u.UserId,
		"email":        u.Email,
		"display_name": u.DisplayName,
		"phone":        u.PhoneNumber,
	}
}

func (u *CreateUser) BeforeCreate() error {
	var Zap *zap.SugaredLogger
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(password)
	if err != nil {
		Zap.Error("Error decrypting plain password to hash", err.Error())
		return err
	}
	return nil
}
