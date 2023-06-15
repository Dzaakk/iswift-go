//go:build wireinject
// +build wireinject

package dashboard

import (
	adminRepository "iswift-go-project/internal/admin/repository"
	adminUsecase "iswift-go-project/internal/admin/usecase"
	cartRepository "iswift-go-project/internal/cart/repository"
	cartUseCase "iswift-go-project/internal/cart/usecase"
	handler "iswift-go-project/internal/dashboard/delivery/http"
	usecase "iswift-go-project/internal/dashboard/usecase"
	discountRepository "iswift-go-project/internal/discount/repository"
	discountUseCase "iswift-go-project/internal/discount/usecase"
	orderRepository "iswift-go-project/internal/order/repository"
	orderUsecase "iswift-go-project/internal/order/usecase"
	orderDetailRepository "iswift-go-project/internal/order_detail/repository"
	orderDetailUseCase "iswift-go-project/internal/order_detail/usecase"
	paymentUseCase "iswift-go-project/internal/payment/usecase"
	productRepository "iswift-go-project/internal/product/repository"
	productUseCase "iswift-go-project/internal/product/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUsecase "iswift-go-project/internal/user/usecase"
	fileUpload "iswift-go-project/pkg/fileupload/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.DashboardHandler {
	wire.Build(
		handler.NewDashboardHandler,
		usecase.NewDashboardUseCase,
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
		adminRepository.NewAdminRepository,
		adminUsecase.NewAdminUseCase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
	)

	return &handler.DashboardHandler{}
}
