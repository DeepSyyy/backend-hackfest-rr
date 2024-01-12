package service_interface

import "github.com/DeepSyyy/backend-hackfest-rr/model/data/request"

type ClanService interface {
	CreateClan(clan request.ClanCreateRequest)
}
