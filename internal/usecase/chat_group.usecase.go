package usecase

import (
	"context"
	"go-backend/internal/dto"
)

type ChatGroupUsecase interface {
	FindAll(ctx context.Context, input dto.ChatGroupFindAllInput) (any, error)
}
