package router

import (
	"net/http"

	"github.com/DeepSyyy/backend-hackfest-rr/controller"
	"github.com/DeepSyyy/backend-hackfest-rr/middleware"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"github.com/gin-gonic/gin"
)

func NewRouter(authC *controller.AuthControllerImpl, userRepo repository_interface.UserRepository, userC *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to my service")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("/register", authC.Register)
	authRouter.POST("/login", authC.Login)

	userRouter := router.Group("/users")
	userRouter.GET("", middleware.DeserializeUser(userRepo), userC.GetUsers)

	return service

}
