package controller

import (
	"fmt"
	"net/http"

	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/response"
	service_interface "github.com/DeepSyyy/backend-hackfest-rr/service/interface"
	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	authService service_interface.AuthService
}

func NewAuthController(service service_interface.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{authService: service}
}

func (controller *AuthControllerImpl) Register(c *gin.Context) {
	createUserRequest := request.RegisterRequest{}
	err := c.ShouldBindJSON(&createUserRequest)
	helper_error.ErrorPanic(err)

	error_exist := controller.authService.Register(createUserRequest)

	if error_exist != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Username or email already exist",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Register Success",
		Data:    nil,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := c.ShouldBindJSON(&loginRequest)
	helper_error.ErrorPanic(err)

	token, err_token := controller.authService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, webResponse)
}
