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

type chatGroupUsecase struct {
	chatGroupRepository repository.ChatGroupRepository
}

func NewChatGroupUsecase(chatGroupRepository repository.ChatGroupRepository) usecase.ChatGroupUsecase {
	return &chatGroupUsecase{
		chatGroupRepository: chatGroupRepository,
	}
}

// FindAll implements [usecase.ChatGroupUsecase].
func (a *chatGroupUsecase) FindAll(ctx context.Context, input dto.ChatGroupFindAllInput) (any, error) {
	data, err := a.chatGroupRepository.GetAll(ctx, input.Query, input.ChatGroupFindAllFilters)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// totalItem: tổng số lượng item
	totalItem, err := a.chatGroupRepository.Count(ctx, input.ChatGroupFindAllFilters)
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
