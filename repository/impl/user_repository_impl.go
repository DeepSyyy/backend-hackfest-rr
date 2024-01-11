package repository_impl

import (
	"errors"

	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/domain"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) repository_interface.UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete
func (u *UserRepositoryImpl) Delete(userId int) {
	var user domain.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper_error.ErrorPanic(result.Error)
}

// FindAll
func (u *UserRepositoryImpl) FindAll() []domain.User {
	var User []domain.User
	result := u.Db.Find(&User)
	helper_error.ErrorPanic(result.Error)

	return User
}

// FindById
func (u *UserRepositoryImpl) FindById(userId int) (domain.User, error) {
	var User domain.User
	result := u.Db.Find(&User, userId)
	if result != nil {
		return User, nil
	} else {
		return User, errors.New("user not found")
	}
}

// findByUsername
func (u *UserRepositoryImpl) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	result := u.Db.Where("username = ?", username).Find(&user)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

// Save
func (u *UserRepositoryImpl) Save(user domain.User) error {
	result := u.Db.Create(&user)
	if result.Error != nil {
		return errors.New("user already exists")
	}
	return nil
}

// Update
func (u *UserRepositoryImpl) Update(user domain.User) {
	var updateUser = request.UpdateUsersRequest{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Id:       user.Id,
	}

	result := u.Db.Model(&user).Updates(updateUser)
	helper_error.ErrorPanic(result.Error)
}
