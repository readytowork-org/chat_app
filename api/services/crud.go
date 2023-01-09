package services

import (
	"letschat/infrastructure"

	"firebase.google.com/go/v4/auth"
)

type CrudService struct {
	fbAuth *auth.Client
	logger infrastructure.Logger
}

func NewCrudService(
	fbAuth *auth.Client,
	logger infrastructure.Logger,
) CrudService {
	return CrudService{
		fbAuth: fbAuth,
		logger: logger,
	}
}

func (cs CrudService) CreateData() {}
func (cs CrudService) UpdateData() {}
func (cs CrudService) DeleteData() {}
func (cs CrudService) GetData()    {}
