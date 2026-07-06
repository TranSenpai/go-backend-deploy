package repository

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"time"
)

type ChatMessageRepository interface {
	FindAll(ctx context.Context) (any, error)
	CreateMessage(ctx context.Context, userId int, chatGroupId int, messageText string, createdAt time.Time) (*ent.ChatMessages, error)
	GetAll(ctx context.Context, query pagination.Query, filters dto.ChatMessageFindAllFilters) (any, error)
	Count(ctx context.Context, filters dto.ChatMessageFindAllFilters) (int, error)
}
