package domain

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
}
