package repository

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
)

type ArticleRepository interface {
	Create(ctx context.Context, body dto.ArticleCreateReq) (any, error)
	CreateGorm(ctx context.Context, body dto.ArticleCreateReq) (any, error)
	GetAll(ctx context.Context, query pagination.Query, filters dto.ArticleFindAllFilters) ([]*ent.Articles, error)
	Count(ctx context.Context, filters dto.ArticleFindAllFilters) (int, error)
	Delete(ctx context.Context, id int) error
}
