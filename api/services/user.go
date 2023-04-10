package services

import (
	"fmt"
	"letschat/api/helper"
	"letschat/api/repository"
	"letschat/infrastructure"
	"letschat/models"
)

type UserService struct {
	userrepository repository.UserRepository
	logger         infrastructure.Logger
}

func NewUserService(
	userrepository repository.UserRepository,
	logger infrastructure.Logger,
) UserService {
	return UserService{
		userrepository: userrepository,
		logger:         logger,
	}
}

func (c UserService) Create(user models.CreateUser) error {
	return c.userrepository.Create(user)
}

func (c UserService) Update(id string, user models.UpdateUser) error {
	return c.userrepository.Update(id, user)
}

func (c UserService) Delete(id string) error {
	return c.userrepository.Delete(id)
}
func (c UserService) FindOne(id string) (*models.User, error) {
	return c.userrepository.FindOne(id)
}

func (c UserService) CheckUserWithPhone(phone string) (*models.CreateUser, bool, error) {
	return c.userrepository.CheckUserWithPhone(phone)
}
func (c UserService) FindAll() (*[]models.User, error) {
	return c.userrepository.FindAll()
}

func (c UserService) UpdateUserStatus(userID string, status bool) error {
	return c.userrepository.UpdateUserStatus(userID, status)
}
func (c UserService) GetAllRooms(userId string) ([]string, error) {
	return c.userrepository.GetAllRooms(userId)
}
func (c UserService) IsRoomPresent(userID string, roomID string) (bool, error) {
	rooms, err := c.userrepository.GetAllRooms(userID)
	if err != nil {
		return false, err
	}
	for _, r := range rooms {
		if r == roomID {
			return true, nil
		}
	}
	return false, nil
}

func (c UserService) AddRoom(userID string, roomId string) error {
	rooms, err := c.userrepository.GetAllRooms(userID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rooms = append(rooms, roomId)
	err = c.userrepository.UpdateRoom(userID, rooms)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c UserService) DeleteRoom(userID string, roomId string) error {
	rooms, err := c.userrepository.GetAllRooms(userID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rooms = helper.RemoveString(rooms, roomId)
	err = c.userrepository.UpdateRoom(userID, rooms)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
