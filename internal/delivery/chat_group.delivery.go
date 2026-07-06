package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type chatGroupDelivery struct {
	chatGroupHandler *handler.ChatGroupHandler
}

func NewChatGroupDelivery(chatGroupHandler *handler.ChatGroupHandler) *chatGroupDelivery {
	return &chatGroupDelivery{
		chatGroupHandler: chatGroupHandler,
	}
}

func (d *chatGroupDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {
	chatGroupGroup := apiGroup.Group("chat-group")
	{
		chatGroupGroup.GET("", d.chatGroupHandler.FindAll)
	}
}
