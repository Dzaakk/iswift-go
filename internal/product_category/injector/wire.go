//go:build wireinject
// +build wireinject

package product_categories

import (
	handler "iswift-go-project/internal/product_category/delivery/http"
	repository "iswift-go-project/internal/product_category/repository"
	usecase "iswift-go-project/internal/product_category/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
		handler.NewProductCategoryHandler,
	)

	return &handler.ProductCategoryHandler{}
}
