package handler

import (
	"fmt"
	"go-backend/internal/common/middlewares"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoUsecase usecase.DemoUsecase
}

func NewDemoHandler(demoUsecase usecase.DemoUsecase) *DemoHandler {
	return &DemoHandler{
		demoUsecase: demoUsecase,
	}
}

func (d *DemoHandler) Query(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	input := pagination.Query{
		Page:     page,
		PageSize: pageSize,
	}

	// tham số nhận vào
	data := d.demoUsecase.Query(input)
	response.Success(data, "", 0, ctx)
}

func (d *DemoHandler) Param(ctx *gin.Context) {
	fmt.Println("")
	fmt.Println("Handler Param (mid)")

	userRaw, exists := ctx.Get("user")
	if !exists {
		fmt.Println("Lỗi không có user")
	}
	user, ok := userRaw.(middlewares.User)
	if !ok {
		fmt.Println("Lỗi không phải struct User")
	}
	fmt.Println("Handler nhận được", user)

	fmt.Println("")

	id, err := strconv.Atoi(ctx.Param("id"))
	// panic("Lỗi không kiểm soát được")
	if err != nil {
		ctx.Error(response.NewBadRequestException("Mật khẩu sai rồi"))
		return
	}

	data := d.demoUsecase.Param(id)
	response.Success(data, "", 0, ctx)
}

func (d *DemoHandler) Body(ctx *gin.Context) {
	var body dto.DemoBody

	err := ctx.ShouldBindJSON(&body)
	fmt.Println("err", err)

	data := d.demoUsecase.Body(body)
	response.Success(data, "", 0, ctx)
}

func (d *DemoHandler) Header(ctx *gin.Context) {
	apiKey := ctx.GetHeader("API_KEY")
	data := d.demoUsecase.Header(apiKey)
	response.Success(data, "", 0, ctx)
}
