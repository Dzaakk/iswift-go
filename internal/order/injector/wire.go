//go:build wireinject
// +build wireinject

package order

import (
	cartRepository "iswift-go-project/internal/cart/repository"
	cartUseCase "iswift-go-project/internal/cart/usecase"
	discountRepository "iswift-go-project/internal/discount/repository"
	discountUseCase "iswift-go-project/internal/discount/usecase"
	handler "iswift-go-project/internal/order/delivery/http"
	repository "iswift-go-project/internal/order/repository"
	usecase "iswift-go-project/internal/order/usecase"
	orderDetailRepository "iswift-go-project/internal/order_detail/repository"
	orderDetailUseCase "iswift-go-project/internal/order_detail/usecase"
	paymentUseCase "iswift-go-project/internal/payment/usecase"
	productRepository "iswift-go-project/internal/product/repository"
	productUseCase "iswift-go-project/internal/product/usecase"
	fileUpload "iswift-go-project/pkg/fileupload/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUsecase,
		fileUpload.NewFileUpload,
	)

	return &handler.OrderHandler{}
}
