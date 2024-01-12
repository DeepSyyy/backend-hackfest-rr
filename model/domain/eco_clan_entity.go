package domain

import (
	"time"

	"github.com/google/uuid"
)

type EcoClan struct {
	ClanTag      string    `gorm:"primaryKey;type:string;not null"`
	Name         string    `gorm:"type:varchar(50);not null;uniqueIndex"`
	Description  string    `gorm:"type:varchar(255);not null"`
	ClanLeaderId uuid.UUID `gorm:"type:uuid;foreignKey:ClanLeaderId"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime:true"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime:true"`
}
