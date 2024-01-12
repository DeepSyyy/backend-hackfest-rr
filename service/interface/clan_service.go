package service_interface

import (
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/response"
)

type ClanService interface {
	CreateClan(clan request.ClanCreateRequest)
	GetClans() []response.ClanResponse
}
