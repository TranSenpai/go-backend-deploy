package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/ent/users"
	"go-backend/internal/common/pagination"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
)

type userRepository struct {
	entClient *ent.Client
}

func NewUserRepository(entClient *ent.Client) repository.UserRepository {
	return &userRepository{
		entClient: entClient,
	}
}

// FindAll implements [repository.UserRepository].
func (a *userRepository) ExitsByEmail(ctx context.Context, email string) (bool, error) {
	entQuery := a.entClient.Users.Query()

	entQuery = entQuery.Where(users.EmailEQ(email))

	return entQuery.Exist(ctx)
}

// CreateUser implements [repository.UserRepository].
func (a *userRepository) CreateUser(ctx context.Context, body dto.AuthRegisterReq) (*ent.Users, error) {
	entCreate := a.entClient.Users.Create()
	entCreate = entCreate.SetEmail(body.Email)
	entCreate = entCreate.SetPassword(body.Password)
	entCreate = entCreate.SetFullName(body.FullName)

	return entCreate.Save(ctx)
}

// FindUserByEmail implements [repository.UserRepository].
func (a *userRepository) FindUserByEmail(ctx context.Context, email string) (*ent.Users, error) {
	entQuery := a.entClient.Users.Query()
	entQuery = entQuery.Where(users.EmailEQ(email))
	return entQuery.Only(ctx)
}

// FindUserById implements [repository.UserRepository].
func (a *userRepository) FindUserById(ctx context.Context, id int) (*ent.Users, error) {
	entQuery := a.entClient.Users.Query()
	entQuery = entQuery.Where(users.IDEQ(id))
	return entQuery.Only(ctx)
}

// CreateUserForGoogle implements [repository.UserRepository].
func (a *userRepository) CreateUserForGoogle(ctx context.Context, body dto.AuthCreateUserForGoogleReq) (*ent.Users, error) {
	entCreate := a.entClient.Users.Create()
	entCreate = entCreate.SetEmail(body.Email)
	entCreate = entCreate.SetFullName(body.FullName)
	entCreate = entCreate.SetAvatar(body.Avatar)
	entCreate = entCreate.SetGoogleID(body.GoogleId)

	return entCreate.Save(ctx)
}

// UploadAvatarById implements [repository.UserRepository].
func (a *userRepository) UpdateAvatarById(ctx context.Context, id int, avatar string) (*ent.Users, error) {
	entUpdate := a.entClient.Users.UpdateOneID(id)
	entUpdate = entUpdate.SetAvatar(avatar)
	return entUpdate.Save(ctx)
}

// GetAll implements [repository.UserRepository].
func (a *userRepository) GetAll(ctx context.Context, query pagination.Query, filters dto.UserFindAllFilters) ([]*ent.Users, error) {
	entQuery := a.entClient.Users.Query()

	handlerFilterUser(filters, entQuery)

	entQuery = entQuery.Limit(query.PageSize)
	entQuery = entQuery.Offset(query.Offset)
	return entQuery.All(ctx)
}

// Count implements [repository.UserRepository].
func (a *userRepository) Count(ctx context.Context, filters dto.UserFindAllFilters) (int, error) {
	entQuery := a.entClient.Users.Query()

	handlerFilterUser(filters, entQuery)

	return entQuery.Count(ctx)
}

func handlerFilterUser(filters dto.UserFindAllFilters, entQuery *ent.UsersQuery) {
	if filters.Name != "" {
		entQuery = entQuery.Where(users.FullNameContainsFold(filters.Name))
	}
}
