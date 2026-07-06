package delivery

import (
	"go-backend/internal/common/middlewares"
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userHandler    *handler.UserHandler
	authMiddleware *middlewares.AuthMiddleware
}

func NewUserDelivery(userHandler *handler.UserHandler, authMiddleware *middlewares.AuthMiddleware) *userDelivery {
	return &userDelivery{
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

func (d *userDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	userGroup := apiGroup.Group("user")
	{
		userGroup.Use(d.authMiddleware.Protect)
		userGroup.GET("", d.userHandler.FindAll)
		userGroup.GET(":id", d.userHandler.FindOne)
		userGroup.POST("avatar-local", d.userHandler.AvatarLocal)
		userGroup.POST("avatar-cloud", d.userHandler.AvatarCloud)
	}
}
