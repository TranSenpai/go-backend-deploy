package handler

import (
	"go-backend/internal/common/response"
	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchUsecase usecase.SearchUsecase
}

func NewSearchHandler(searchUsecase usecase.SearchUsecase) *SearchHandler {
	return &SearchHandler{
		searchUsecase: searchUsecase,
	}
}

func (a *SearchHandler) FindAll(ctx *gin.Context) {
	textSearch := ctx.Query("text")

	result, err := a.searchUsecase.FindAll(ctx.Request.Context(), textSearch)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
