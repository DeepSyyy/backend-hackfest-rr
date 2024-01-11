package service_impl

import (
	"errors"

	"github.com/DeepSyyy/backend-hackfest-rr/config"
	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/domain"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	service_interface "github.com/DeepSyyy/backend-hackfest-rr/service/interface"
	"github.com/DeepSyyy/backend-hackfest-rr/utils"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRepository repository_interface.UserRepository
	Validate       *validator.Validate
}

func NewAuthService(userRepository repository_interface.UserRepository, validate *validator.Validate) service_interface.AuthService {
	return &AuthServiceImpl{UserRepository: userRepository, Validate: validate}
}

// Login UseCase
func (a *AuthServiceImpl) Login(user request.LoginRequest) (string, error) {
	new_user, user_err := a.UserRepository.FindByUsername(user.Username)
	if user_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_user.Password, user.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.TokenSecret)
	helper_error.ErrorPanic(err_token)
	return token, nil
}

// Register UseCase
func (a *AuthServiceImpl) Register(user request.RegisterRequest) error {
	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	helper_error.ErrorPanic(err)

	newUser := domain.User{
		Username: user.Username,
		Password: hashedPassword,
		Email:    user.Email,
	}

	error_exist := a.UserRepository.Save(newUser)

	return error_exist
}
