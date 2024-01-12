package controller

import (
	"net/http"

	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/response"
	service_interface "github.com/DeepSyyy/backend-hackfest-rr/service/interface"
	"github.com/gin-gonic/gin"
)

type ClanController struct {
	clanService service_interface.ClanService
}

func NewClanController(clanService service_interface.ClanService) *ClanController {
	return &ClanController{clanService: clanService}
}

// CreateClan
func (c *ClanController) CreateClan(ctx *gin.Context) {
	createClanRequest := request.ClanCreateRequest{}
	err := ctx.ShouldBindJSON(&createClanRequest)
	helper_error.ErrorPanic(err)

	c.clanService.CreateClan(createClanRequest)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully create clan!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
