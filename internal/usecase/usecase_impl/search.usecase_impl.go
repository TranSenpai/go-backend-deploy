package usecase_impl

import (
	"context"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
)

type searchUsecase struct {
	searchRepository repository.SearchRepository
}

func NewSearchUsecase(searchRepository repository.SearchRepository) usecase.SearchUsecase {
	return &searchUsecase{
		searchRepository: searchRepository,
	}
}

// FindAll implements [usecase.SearchUsecase].
func (a *searchUsecase) FindAll(ctx context.Context, textSearch string) (any, error) {
	return a.searchRepository.FindAll(ctx, textSearch)
}
