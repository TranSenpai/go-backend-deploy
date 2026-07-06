package usecase

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/dto"
	"mime/multipart"
)

type UserUsecase interface {
	FindAll(ctx context.Context, input dto.UserFindAllInput) (any, error)
	FindOne(ctx context.Context, id int) (any, error)
	AvatarLocal(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error)
	AvatarCloud(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error)
}
