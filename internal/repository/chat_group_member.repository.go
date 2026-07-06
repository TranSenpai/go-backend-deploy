package repository

import (
	"context"
)

type ChatGroupMemberRepository interface {
	FindAll(ctx context.Context) (any, error)
	CreateGroupMemberMany(ctx context.Context, chatGroupId int, ids []int) error
}
