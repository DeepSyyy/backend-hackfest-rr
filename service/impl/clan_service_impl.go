package service_impl

import (
	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/data/request"
	"github.com/DeepSyyy/backend-hackfest-rr/model/domain"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	service_interface "github.com/DeepSyyy/backend-hackfest-rr/service/interface"
	"github.com/DeepSyyy/backend-hackfest-rr/utils"
	"github.com/go-playground/validator/v10"
)

type ClanServiceImpl struct {
	ClanRepository repository_interface.ClanRepository
	Validate       *validator.Validate
}

func NewClanServiceImpl(ClanRepository repository_interface.ClanRepository, Validate *validator.Validate) service_interface.ClanService {
	return &ClanServiceImpl{ClanRepository: ClanRepository, Validate: Validate}
}

// CreateClan
func (c *ClanServiceImpl) CreateClan(clan request.ClanCreateRequest) {
	newClan := domain.EcoClan{
		ClanTag:     utils.GenerateRandomClanTag(),
		Name:        clan.Name,
		Description: clan.Description,
	}

	err := c.ClanRepository.Save(newClan)
	helper_error.ErrorPanic(err)
}
