//go:build wireinject
// +build wireinject

package cart

import (
	handler "iswift-go-project/internal/cart/delivery/http"
	repository "iswift-go-project/internal/cart/repository"
	usecase "iswift-go-project/internal/cart/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		handler.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
	)

	return &handler.CartHandler{}
}
