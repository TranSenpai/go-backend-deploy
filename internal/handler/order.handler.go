package handler

import (
	"go-backend/internal/common/helpers"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"io"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

func (a *OrderHandler) FindAll(ctx *gin.Context) {
	result, err := a.orderUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}

func (a *OrderHandler) Create(ctx *gin.Context) {
	var body dto.CreateOrder
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		if err == io.EOF {
			ctx.Error(response.NewBadRequestException("Body required"))
			return
		}
		ctx.Error(response.NewBadRequestException(err.Error()))
		return
	}

	user, err := helpers.GetUser(ctx)
	if err != nil {
		ctx.Error(response.NewUnauthorizedException())
		return
	}
	body.UserId = user.ID

	result, err := a.orderUsecase.CreateRequest(ctx.Request.Context(), body)
	if err != nil {
		ctx.Error(err)
		return
	}

	response.Success(result, "", 0, ctx)
}
