package repository

import "demoG/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (users model.Users, err error)
	FindAll() []model.Users
}
