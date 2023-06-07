//go:build wireinject
// +build wireinject

package discount

import (
	handler "iswift-go-project/internal/discount/delivery/http"
	repository "iswift-go-project/internal/discount/repository"
	usecase "iswift-go-project/internal/discount/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.DiscountHandler {
	wire.Build(
		repository.NewDiscountRepository,
		usecase.NewDiscountUseCase,
		handler.NewDiscountHandler,
	)

	return &handler.DiscountHandler{}
}
