package handler

import (
	"encoding/json"
	"fmt"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(articleUsecase usecase.ArticleUsecase) *ArticleHandler {
	return &ArticleHandler{
		articleUsecase: articleUsecase,
	}
}

func (a *ArticleHandler) Create(ctx *gin.Context) {
	var body dto.ArticleCreateReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	reuslt, err := a.articleUsecase.Create(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}

func (a *ArticleHandler) FindAll(ctx *gin.Context) {
	queryPagi := pagination.Get(
		ctx.Query("page"),
		ctx.Query("pageSize"),
	)

	filterString := ctx.DefaultQuery("filters", "{}")
	var filters dto.ArticleFindAllFilters
	json.Unmarshal([]byte(filterString), &filters)

	fmt.Printf("%+v \n\n", filters)

	input := dto.ArticleFindAllInput{
		Query:                 queryPagi,
		ArticleFindAllFilters: filters,
	}

	fmt.Printf("%+v \n\n", queryPagi)

	reuslt, err := a.articleUsecase.FindAll(ctx.Request.Context(), input)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}

func (a *ArticleHandler) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	reuslt, err := a.articleUsecase.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(reuslt, "", 0, ctx)
}
