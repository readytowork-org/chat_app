package services

import (
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

func (c UserService) Create(user models.User) error {
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
func (c UserService) FindAll() (*[]models.User, error) {
	return c.userrepository.FindAll()
}
