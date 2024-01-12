package controller

import (
	"net/http"

	"github.com/DeepSyyy/backend-hackfest-rr/model/data/response"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository_interface.UserRepository
}

func NewUsersController(repository repository_interface.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(model.Users)
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
