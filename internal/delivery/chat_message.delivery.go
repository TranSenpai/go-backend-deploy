package delivery

import (
	"go-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

type chatMessageDelivery struct {
	chatMessageHandler *handler.ChatMessageHandler
}

func NewChatMessageDelivery(chatMessageHandler *handler.ChatMessageHandler) *chatMessageDelivery {
	return &chatMessageDelivery{
		chatMessageHandler: chatMessageHandler,
	}
}

func (d *chatMessageDelivery) RegisterRouter(apiGroup *gin.RouterGroup) {

	chatMessageGroup := apiGroup.Group("chat-message")
	{
		chatMessageGroup.GET("", d.chatMessageHandler.FindAll)
	}
}
