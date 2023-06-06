//go:build wireinject
// +build wireinject

package product

import (
	handler "iswift-go-project/internal/product/delivery/http"
	repository "iswift-go-project/internal/product/repository"
	usecase "iswift-go-project/internal/product/usecase"
	fileUpload "iswift-go-project/pkg/fileupload/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductHanlder {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUsecase,
		repository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &handler.ProductHanlder{}
}
