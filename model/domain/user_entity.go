package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"uniqueIndex;not null"`
}
