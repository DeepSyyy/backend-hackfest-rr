package repository_interface

import "github.com/DeepSyyy/backend-hackfest-rr/model/domain"

type UserRepository interface {
	Save(user domain.User) error
	FindByUsername(username string) (domain.User, error)
	FindAll() []domain.User
	FindById(userId int) (domain.User, error)
	Update(user domain.User)
	Delete(userid int)
}
