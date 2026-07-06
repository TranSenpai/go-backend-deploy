package usecase_impl

import (
	"context"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
	"math"
)

type articleUsecase struct {
	articleRepository repository.ArticleRepository
}

func NewArticleUsecase(articleRepository repository.ArticleRepository) usecase.ArticleUsecase {
	return &articleUsecase{
		articleRepository: articleRepository,
	}
}

// Create implements [usecase.ArticleUsecase].
func (a *articleUsecase) Create(ctx context.Context, body dto.ArticleCreateReq) (any, error) {
	data, err := a.articleRepository.Create(ctx, body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	return data, nil
}

// FindAll implements [usecase.ArticleUsecase].
func (a *articleUsecase) FindAll(ctx context.Context, input dto.ArticleFindAllInput) (any, error) {
	data, err := a.articleRepository.GetAll(ctx, input.Query, input.ArticleFindAllFilters)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// totalItem: tổng số lượng item
	totalItem, err := a.articleRepository.Count(ctx, input.ArticleFindAllFilters)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// totalPage: tổng số trang totalItem / pageSize
	totalPage := float64(totalItem) / float64(input.PageSize)

	result := pagination.PaginationRes[any]{
		Items:     data,
		Page:      input.Page,
		PageSize:  input.PageSize,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(totalPage)),
	}

	return result, nil
}

// Delete implements [usecase.ArticleUsecase].
func (a *articleUsecase) Delete(ctx context.Context, id int) (any, error) {
	err := a.articleRepository.Delete(ctx, id)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return true, nil
}
