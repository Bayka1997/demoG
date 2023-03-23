package service

import (
	"demoG/data/request"
	"demoG/data/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(userId int)
	FindById(userId int) response.UsersResponse
	FindAll() []response.UsersResponse
}
