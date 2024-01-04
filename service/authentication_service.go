package service

import (

	"github.com/prathishbv/notes-api/data/request"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest)
}