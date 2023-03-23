package service

import (
	"demoG/data/request"
	"demoG/data/response"
	"demoG/helper"
	"demoG/model"
	"demoG/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements UsersService
func (t *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := model.Users{
		Name: users.Name,
	}
	t.UsersRepository.Save(userModel)
}

// Delete implements UsersService
func (t *UsersServiceImpl) Delete(usersId int) {
	t.UsersRepository.Delete(usersId)
}

// FindAll implements UsersService
func (t *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := t.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		users = append(users, user)
	}

	return users
}

// FindById implements UsersService
func (t *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	userData, err := t.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:   userData.Id,
		Name: userData.Name,
	}
	return userResponse
}

// Update implements UsersService
func (t *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
	userData, err := t.UsersRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	t.UsersRepository.Update(userData)
}
