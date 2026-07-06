package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/ent/articles"
	"go-backend/internal/common/models"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/repository"

	"gorm.io/gorm"
)

type articleRepository struct {
	entClient  *ent.Client
	gormClient *gorm.DB
}

func NewArticleRepository(entClient *ent.Client, gormClient *gorm.DB) repository.ArticleRepository {
	return &articleRepository{
		entClient:  entClient,
		gormClient: gormClient,
	}
}

// Create implements [repository.ArticleRepository].
func (a *articleRepository) Create(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	entCreate := a.entClient.Articles.Create()

	entCreate = entCreate.SetTitle(body.Title)

	if body.Content != nil {
		entCreate = entCreate.SetContent(*body.Content)
	}

	if body.ImageUrl != nil {
		entCreate = entCreate.SetImageURL(*body.ImageUrl)
	}

	entCreate = entCreate.SetUserID(1)

	return entCreate.Save(ctx)

}

// CreateGorm implements [repository.ArticleRepository].
func (a *articleRepository) CreateGorm(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	article := models.Article{
		Title:    body.Title,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
	}

	result := gorm.WithResult()
	err := gorm.G[models.Article](a.gormClient, result).Create(ctx, &article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements [repository.ArticleRepository].
func (a *articleRepository) GetAll(ctx context.Context, query pagination.Query, filters dto.ArticleFindAllFilters) ([]*ent.Articles, error) {
	entQuery := a.entClient.Articles.Query()

	handlerFilter(filters, entQuery)

	entQuery = entQuery.WithUsers()

	entQuery = entQuery.Limit(query.PageSize)
	entQuery = entQuery.Offset(query.Offset)
	return entQuery.All(ctx)
}

// Count implements [repository.ArticleRepository].
func (a *articleRepository) Count(ctx context.Context, filters dto.ArticleFindAllFilters) (int, error) {
	entQuery := a.entClient.Articles.Query()

	handlerFilter(filters, entQuery)

	return entQuery.Count(ctx)
}

func handlerFilter(filters dto.ArticleFindAllFilters, entQuery *ent.ArticlesQuery) {
	if filters.Id > 0 {
		entQuery = entQuery.Where(articles.IDEQ(filters.Id))
	}

	if filters.Content != "" {
		entQuery = entQuery.Where(articles.ContentContainsFold(filters.Content))
	}

	if filters.Views != nil {
		entQuery = entQuery.Where(articles.ViewsEQ(*filters.Views))
	}
}

// Delete implements [repository.ArticleRepository].
func (a *articleRepository) Delete(ctx context.Context, id int) error {
	entDelete := a.entClient.Articles.DeleteOneID(id)
	return entDelete.Exec(ctx)
}
