package repository_interface

import "github.com/DeepSyyy/backend-hackfest-rr/model/domain"

type ClanRepository interface {
	Save(clan domain.EcoClan) error
	FindAll() []domain.EcoClan
	// FindById(clanId string) (domain.EcoClan, error)
	// Update(clan domain.EcoClan)
	// Delete(clanId string)
}
