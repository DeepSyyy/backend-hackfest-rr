package router

import (
	"net/http"

	"github.com/DeepSyyy/backend-hackfest-rr/controller"
	"github.com/DeepSyyy/backend-hackfest-rr/middleware"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"github.com/gin-gonic/gin"
)

func NewRouter(authC *controller.AuthControllerImpl, userRepo repository_interface.UserRepository, userC *controller.UserController, clanC *controller.ClanController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to my service")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	//Auth Router
	authRouter := router.Group("/auth")
	authRouter.POST("/register", authC.Register)
	authRouter.POST("/login", authC.Login)

	//User Router
	userRouter := router.Group("/users")
	userRouter.GET("", middleware.DeserializeUser(userRepo), userC.GetUsers)

	//Clan Router
	clanRouter := router.Group("/clans")
	clanRouter.POST("", middleware.DeserializeUser(userRepo), clanC.CreateClan)

	return service

}
