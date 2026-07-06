package usecase_impl

import (
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
)

type DemoUsecase struct{}

func NewDemoUsecase() usecase.DemoUsecase {
	return &DemoUsecase{}
}

func (d *DemoUsecase) Query(input pagination.Query) any {
	return input
}

func (d *DemoUsecase) Param(id int) int {
	return id
}

func (d *DemoUsecase) Body(body dto.DemoBody) any {
	return body
}

func (d *DemoUsecase) Header(apiKey string) any {
	return apiKey
}
