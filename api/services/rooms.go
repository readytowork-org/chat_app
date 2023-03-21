package services

import (
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

func (c RoomService) Create(room models.Room) error {
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
