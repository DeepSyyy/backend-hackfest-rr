package service_interface

import "github.com/DeepSyyy/backend-hackfest-rr/model/data/request"

type AuthService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.RegisterRequest) error
}
