package repository

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
)

type ChatGroupRepository interface {
	FindAll(ctx context.Context) (any, error)
	CheckChatGroupOneOneExist(ctx context.Context, ids []int) (*ent.ChatGroups, error)
	FindOneById(ctx context.Context, id int) (*ent.ChatGroups, error)
	CreateGroup(ctx context.Context, name string, userId int) (*ent.ChatGroups, error)
	GetAll(ctx context.Context, query pagination.Query, filters dto.ChatGroupFindAllFilters) (any, error)
	Count(ctx context.Context, filters dto.ChatGroupFindAllFilters) (int, error)
}
