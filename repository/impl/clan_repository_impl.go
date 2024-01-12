package repository_impl

import (
	"github.com/DeepSyyy/backend-hackfest-rr/model/domain"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"gorm.io/gorm"
)

type ClanRepositoryImpl struct {
	Db *gorm.DB
}

func NewClanRepositoryImpl(Db *gorm.DB) repository_interface.ClanRepository {
	return &ClanRepositoryImpl{Db: Db}
}

// Save
func (c *ClanRepositoryImpl) Save(clan domain.EcoClan) error {
	result := c.Db.Create(&clan)
	print(result.Error)
	return result.Error
}

// find all
func (c *ClanRepositoryImpl) FindAll() []domain.EcoClan {
	var clan []domain.EcoClan
	result := c.Db.Find(&clan)
	if result != nil {
		return clan
	} else {
		return clan
	}
}
