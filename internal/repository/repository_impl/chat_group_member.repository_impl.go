package repository_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/repository"
)

type chatGroupMemberRepository struct {
	entClient *ent.Client
}

func NewChatGroupMemberRepository(entClient *ent.Client) repository.ChatGroupMemberRepository {
	return &chatGroupMemberRepository{
		entClient: entClient,
	}
}

// FindAll implements [repository.ChatGroupMemberRepository].
func (a *chatGroupMemberRepository) FindAll(ctx context.Context) (any, error) {
	return nil, nil
}

// CreateGroupMemberMany implements [repository.ChatGroupMemberRepository].
func (a *chatGroupMemberRepository) CreateGroupMemberMany(ctx context.Context, chatGroupId int, ids []int) error {
	entClientTx := GetClientTx(ctx, a.entClient)
	builders := []*ent.ChatGroupMembersCreate{}
	for _, id := range ids {
		builders = append(
			builders,
			entClientTx.ChatGroupMembers.Create().SetUserID(id).SetChatGroupID(chatGroupId),
		)
	}
	return entClientTx.ChatGroupMembers.CreateBulk(builders...).Exec(ctx)
}
