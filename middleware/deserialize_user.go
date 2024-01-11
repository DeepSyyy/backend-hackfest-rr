package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DeepSyyy/backend-hackfest-rr/config"
	helper_error "github.com/DeepSyyy/backend-hackfest-rr/helper/error"
	repository_interface "github.com/DeepSyyy/backend-hackfest-rr/repository/interface"
	"github.com/DeepSyyy/backend-hackfest-rr/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository_interface.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper_error.ErrorPanic(err_id)
		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Next()

	}
}
