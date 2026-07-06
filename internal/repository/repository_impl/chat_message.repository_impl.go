package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/ent/chatmessages"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"time"
)

type chatMessageRepository struct {
	entClient *ent.Client
}

func NewChatMessageRepository(entClient *ent.Client) repository.ChatMessageRepository {
	return &chatMessageRepository{
		entClient: entClient,
	}
}

// FindAll implements [repository.ChatMessageRepository].
func (a *chatMessageRepository) FindAll(ctx context.Context) (any, error) {
	return nil, nil
}

// CreateMessage implements [repository.ChatMessageRepository].
func (a *chatMessageRepository) CreateMessage(ctx context.Context, userId int, chatGroupId int, messageText string, createdAt time.Time) (*ent.ChatMessages, error) {
	entCreate := a.entClient.ChatMessages.Create()
	entCreate = entCreate.SetUserID(userId).
		SetChatGroupID(chatGroupId).
		SetMessageText(messageText).
		SetCreatedAt(createdAt)

	return entCreate.Save(ctx)
}

// GetAll implements [repository.ChatMessageRepository].
func (a *chatMessageRepository) GetAll(ctx context.Context, query pagination.Query, filters dto.ChatMessageFindAllFilters) (any, error) {
	entQuery := a.entClient.ChatMessages.Query()

	handlerFilterChatMessage(filters, entQuery)

	entQuery = entQuery.WithUsers()

	entQuery = entQuery.Limit(query.PageSize)
	entQuery = entQuery.Offset(query.Offset)

	entQuery = entQuery.Order(ent.Desc(chatmessages.FieldCreatedAt))

	return entQuery.All(ctx)
}

// Count implements [repository.ChatMessageRepository].
func (a *chatMessageRepository) Count(ctx context.Context, filters dto.ChatMessageFindAllFilters) (int, error) {
	entQuery := a.entClient.ChatMessages.Query()

	handlerFilterChatMessage(filters, entQuery)

	return entQuery.Count(ctx)
}

func handlerFilterChatMessage(filters dto.ChatMessageFindAllFilters, entQuery *ent.ChatMessagesQuery) {
	if filters.MessageText != "" {
		entQuery = entQuery.Where(chatmessages.MessageTextContainsFold(filters.MessageText))
	}

	entQuery = entQuery.Where(chatmessages.ChatGroupIDEQ(filters.ChatGroupId))
}
