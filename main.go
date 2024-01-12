package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DeepSyyy/backend-hackfest-rr/config"
	"github.com/DeepSyyy/backend-hackfest-rr/controller"
	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	"github.com/DeepSyyy/backend-hackfest-rr/model/domain"
	repository_impl "github.com/DeepSyyy/backend-hackfest-rr/repository/impl"
	"github.com/DeepSyyy/backend-hackfest-rr/router"
	service_impl "github.com/DeepSyyy/backend-hackfest-rr/service/impl"
	"github.com/go-playground/validator/v10"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&domain.User{})
	db.Table("eco_clans").AutoMigrate(&domain.EcoClan{})
	fmt.Println("🚀 Migrated Successfully")

	//init Repo
	userRepository := repository_impl.NewUserRepositoryImpl(db)
	clanRepository := repository_impl.NewClanRepositoryImpl(db)

	//init Service
	authService := service_impl.NewAuthService(userRepository, validate)
	clanService := service_impl.NewClanServiceImpl(clanRepository, validate)

	//init controller
	authController := controller.NewAuthController(authService)
	userController := controller.NewUsersController(userRepository)
	clanController := controller.NewClanController(clanService)

	routes := router.NewRouter(authController, userRepository, userController, clanController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper_error.ErrorPanic(server_err)
}
