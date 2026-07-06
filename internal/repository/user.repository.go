package repository

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
)

type UserRepository interface {
	ExitsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, body dto.AuthRegisterReq) (*ent.Users, error)
	CreateUserForGoogle(ctx context.Context, body dto.AuthCreateUserForGoogleReq) (*ent.Users, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.Users, error)
	FindUserById(ctx context.Context, id int) (*ent.Users, error)
	UpdateAvatarById(ctx context.Context, id int, avatar string) (*ent.Users, error)
	GetAll(ctx context.Context, query pagination.Query, filters dto.UserFindAllFilters) ([]*ent.Users, error)
	Count(ctx context.Context, filters dto.UserFindAllFilters) (int, error)
}
