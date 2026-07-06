package usecase_impl

import (
	"context"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
)

type chatMessageUsecase struct {
	chatMessageRepository repository.ChatMessageRepository
}

func NewChatMessageUsecase(chatMessageRepository repository.ChatMessageRepository) usecase.ChatMessageUsecase {
	return &chatMessageUsecase{
		chatMessageRepository: chatMessageRepository,
	}
}

// FindAll implements [usecase.ChatMessageUsecase].
func (a *chatMessageUsecase) FindAll(ctx context.Context) (any, error) {
	return "FindAll", nil
}
