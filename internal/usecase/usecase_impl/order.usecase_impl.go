package usecase_impl

import (
	"context"
	"go-backend/ent"
	"go-backend/internal/common/rabbitmq"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"go-backend/internal/usecase"
)

type orderUsecase struct {
	rabbitmq *rabbitmq.RabbitMQ
}

func NewOrderUsecase(rabbitmq *rabbitmq.RabbitMQ) usecase.OrderUsecase {
	return &orderUsecase{
		rabbitmq: rabbitmq,
	}
}

// FindAll implements [usecase.OrderUsecase].
func (a *orderUsecase) FindAll(ctx context.Context) (any, error) {
	return "FindAll", nil
}

// CreateSend implements [usecase.OrderUsecase].
func (a *orderUsecase) CreateSend(ctx context.Context, body dto.CreateOrder) (any, error) {
	err := a.rabbitmq.Send(ctx, "CREATE_ORDER_SEND", body)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return true, nil
}

// CreateRequest implements [usecase.OrderUsecase].
func (a *orderUsecase) CreateRequest(ctx context.Context, body dto.CreateOrder) (any, error) {
	var result *ent.Orders

	err := a.rabbitmq.Request(
		ctx,
		"CREATE_ORDER_REQUEST",
		body,
		&result,
	)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	return result, nil
}
