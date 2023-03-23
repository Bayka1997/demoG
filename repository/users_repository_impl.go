package repository

import (
	"demoG/data/request"
	"demoG/helper"
	"demoG/model"
	"errors"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersREpositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (t *UsersRepositoryImpl) Delete(usersId int) {
	var users model.Users
	result := t.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository
func (t *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := t.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository
func (t *UsersRepositoryImpl) FindById(usersId int) (users model.Users, err error) {
	var user model.Users
	result := t.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UsersRepository
func (t *UsersRepositoryImpl) Save(users model.Users) {
	result := t.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (t *UsersRepositoryImpl) Update(user model.Users) {
	var updateUser = request.UpdateUsersRequest{
		Id:   user.Id,
		Name: user.Name,
	}
	result := t.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
