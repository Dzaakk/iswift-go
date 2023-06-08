package order_detail

import (
	"fmt"
	entity "iswift-go-project/internal/order_detail/entity"
	repository "iswift-go-project/internal/order_detail/repository"
)

type OrderDetailUsecase interface {
	Create(entity entity.OrderDetail)
}

type OrderDetailUsecaseImpl struct {
	repository repository.OrderDetailRepository
}

// Create implements OrderDetailUsecase.
func (usecase *OrderDetailUsecaseImpl) Create(entity entity.OrderDetail) {
	_, err := usecase.repository.Create(entity)

	if err != nil {
		fmt.Print(err)
	}
}

func NewOrderDetailUseCase(repository repository.OrderDetailRepository) OrderDetailUsecase {
	return &OrderDetailUsecaseImpl{repository}
}
