package repository

import "github.com/prathishbv/notes-api/model"

type UsersRepository interface {
	Save(users model.Users)
	Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (model.Users, error)
	FindAll() []model.Users
	FindByUsername(username string) (model.Users, error)
}