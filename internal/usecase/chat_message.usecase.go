package usecase

import (
	"context"
)

type ChatMessageUsecase interface {
	FindAll(ctx context.Context) (any, error)
}
