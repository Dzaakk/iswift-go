//go:build wireinject
// +build wireinject

package webhook

import (
	cartRepository "iswift-go-project/internal/cart/repository"
	cartUseCase "iswift-go-project/internal/cart/usecase"
	classRoomRepository "iswift-go-project/internal/class_room/repository"
	classRoomUsecase "iswift-go-project/internal/class_room/usecase"
	discountRepository "iswift-go-project/internal/discount/repository"
	discountUseCase "iswift-go-project/internal/discount/usecase"
	orderRepository "iswift-go-project/internal/order/repository"
	orderUsecase "iswift-go-project/internal/order/usecase"
	orderDetailRepository "iswift-go-project/internal/order_detail/repository"
	orderDetailUseCase "iswift-go-project/internal/order_detail/usecase"
	paymentUseCase "iswift-go-project/internal/payment/usecase"
	productRepository "iswift-go-project/internal/product/repository"
	productUseCase "iswift-go-project/internal/product/usecase"
	handler "iswift-go-project/internal/webhook/delivery/http"
	usecase "iswift-go-project/internal/webhook/usecase"
	fileUpload "iswift-go-project/pkg/fileupload/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.WebhookHandler {
	wire.Build(
		handler.NewWebHookHandler,
		usecase.NewWebhookUseCase,
		classRoomRepository.NewClassRoomRepository,
		classRoomUsecase.NewClassRoomUseCase,
		orderRepository.NewOrderRepository,
		orderUsecase.NewOrderUseCase,
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUsecase,
		fileUpload.NewFileUpload,
	)

	return &handler.WebhookHandler{}
}
