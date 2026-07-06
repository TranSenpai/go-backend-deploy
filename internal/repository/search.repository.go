package repository

import (
	"context"
)

type SearchRepository interface {
	FindAll(ctx context.Context, textSearch string) (any, error)
}
