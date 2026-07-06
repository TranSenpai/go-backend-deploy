package usecase

import (
	"context"
	"go-backend/internal/dto"
)

type ArticleUsecase interface {
	Create(ctx context.Context, body dto.ArticleCreateReq) (any, error)
	FindAll(ctx context.Context, input dto.ArticleFindAllInput) (any, error)
	Delete(ctx context.Context, id int) (any, error)
}
