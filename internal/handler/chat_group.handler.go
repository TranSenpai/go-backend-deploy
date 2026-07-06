package handler

import (
	"encoding/json"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ChatGroupHandler struct {
	chatGroupUsecase usecase.ChatGroupUsecase
}

func NewChatGroupHandler(chatGroupUsecase usecase.ChatGroupUsecase) *ChatGroupHandler {
	return &ChatGroupHandler{
		chatGroupUsecase: chatGroupUsecase,
	}
}

func (a *ChatGroupHandler) FindAll(ctx *gin.Context) {
	queryPagi := pagination.Get(
		ctx.Query("page"),
		ctx.Query("pageSize"),
	)

	filterString := ctx.DefaultQuery("filters", "{}")
	var filters dto.ChatGroupFindAllFilters
	json.Unmarshal([]byte(filterString), &filters)

	input := dto.ChatGroupFindAllInput{
		Query:                   queryPagi,
		ChatGroupFindAllFilters: filters,
	}

	result, err := a.chatGroupUsecase.FindAll(ctx.Request.Context(), input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
