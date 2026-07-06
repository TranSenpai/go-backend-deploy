package handler

import (
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ChatMessageHandler struct {
	chatMessageUsecase usecase.ChatMessageUsecase
}

func NewChatMessageHandler(chatMessageUsecase usecase.ChatMessageUsecase) *ChatMessageHandler {
	return &ChatMessageHandler{
		chatMessageUsecase: chatMessageUsecase,
	}
}

func (a *ChatMessageHandler) FindAll(ctx *gin.Context) {
	result, err := a.chatMessageUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
