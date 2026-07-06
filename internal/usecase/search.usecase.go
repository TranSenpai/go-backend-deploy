package usecase

import (
	"context"
)

type SearchUsecase interface {
	FindAll(ctx context.Context, textSearch string) (any, error)
}
