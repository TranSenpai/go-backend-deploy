package usecase

import (
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
)

type DemoUsecase interface {
	Query(input pagination.Query) any
	Param(id int) int
	Body(body dto.DemoBody) any
	Header(apiKey string) any
}
