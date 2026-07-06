package usecase

import (
	"context"
	"go-backend/internal/dto"
)

type OrderUsecase interface {
	FindAll(ctx context.Context) (any, error)
	CreateSend(ctx context.Context, body dto.CreateOrder) (any, error)
	CreateRequest(ctx context.Context, body dto.CreateOrder) (any, error)
}
