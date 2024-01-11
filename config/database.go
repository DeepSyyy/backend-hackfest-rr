package config

import (
	"fmt"

	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper_error.ErrorPanic(err)

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
