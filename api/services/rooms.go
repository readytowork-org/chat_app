package services

import (
	"fmt"
	"letschat/api/helper"
	"letschat/api/repository"
	"letschat/infrastructure"
	"letschat/models"
)

type RoomService struct {
	roomrepository repository.RoomRepository
	logger         infrastructure.Logger
}

func NewRoomService(
	roomrepository repository.RoomRepository,
	logger infrastructure.Logger,
) RoomService {
	return RoomService{
		roomrepository: roomrepository,
		logger:         logger,
	}
}

func (c RoomService) Create(room models.Room) (models.Room, error) {
	return c.roomrepository.Create(room)
}

func (c RoomService) Update(id string, room models.RoomUpdate) error {
	return c.roomrepository.Update(id, room)
}

func (c RoomService) Delete(id string) error {
	return c.roomrepository.Delete(id)
}

func (c RoomService) FindOne(id string) (*models.Room, error) {
	return c.roomrepository.FindOne(id)
}

func (c RoomService) GetAllMembers(roomId string) ([]string, error) {
	return c.roomrepository.GetAllMembers(roomId)
}

func (c RoomService) AddMember(roomId string, memberId string) error {
	members, err := c.roomrepository.GetAllMembers(roomId)
	if err != nil {
		fmt.Println("get all members ///////////////////")
		fmt.Println(err)
		return err
	}
	members = append(members, memberId)
	err = c.roomrepository.UpdateMembers(roomId, members)
	if err != nil {
		fmt.Println("update members members ///////////////////")
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomService) DeleteMember(roomId string, memberId string) error {
	members, err := c.roomrepository.GetAllMembers(roomId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	members = helper.RemoveString(members, memberId)
	err = c.roomrepository.UpdateMembers(roomId, members)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomService) DeleteAllMember(roomId string) error {
	err := c.roomrepository.UpdateMembers(roomId, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c RoomService) UpdateLastMessage(roomId string, messageId string) error {
	return c.roomrepository.UpdateLastMessage(roomId, messageId)
}
