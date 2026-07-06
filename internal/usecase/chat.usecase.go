package usecase

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/dto"
	"time"
)

type ChatUsecase interface {
	CreateGroup(ctx context.Context, accessToken string, targetUserIds []int, name string) (*ent.ChatGroups, error)
	JoinGroup(ctx context.Context, accessToken string, chatGroupId int) (*ent.ChatGroups, error)
	SendMessage(ctx context.Context, accessToken string, chatGroupId int, message string, createdAt time.Time) (*dto.SendMessageReturn, error)
}
