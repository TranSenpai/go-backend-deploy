package repository_impl

import (
	"context"
	"errors"
	"fmt"
	"go-backend/ent"
	"go-backend/internal/repository"
)

type keyTx struct{}

type UnitOfWorkRepository struct {
	entClient *ent.Client
}

func NewUnitOfWorkRepository(entClient *ent.Client) repository.UnitOfWorkRepository {
	return &UnitOfWorkRepository{
		entClient: entClient,
	}
}

// Do implements [repository.UnitOfWorkRepository].
func (u *UnitOfWorkRepository) Do(ctx context.Context, fn func(ctxTx context.Context) error) (err error) {
	tx, err := u.entClient.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			// rollback
			errorRollBack := tx.Rollback()
			if errorRollBack != nil {
				fmt.Println("errorRollBack", errorRollBack)
				err = errors.New("rolback err")
			}
		} else {
			// commit
			errorCommit := tx.Commit()
			if errorCommit != nil {
				fmt.Println("errorCommit", errorCommit)
				err = errors.New("commit err")
			}
		}
	}()

	ctxTx := context.WithValue(ctx, keyTx{}, tx.Client())

	err = fn(ctxTx)
	if err != nil {
		return err
	}

	return nil
}

func GetClientTx(ctx context.Context, client *ent.Client) *ent.Client {
	clientAny := ctx.Value(keyTx{})

	clientTx, ok := clientAny.(*ent.Client)
	if !ok {
		return client
	}

	return clientTx
}
