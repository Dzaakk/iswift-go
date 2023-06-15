package dashboard

import (
	adminUsecase "iswift-go-project/internal/admin/usecase"
	dto "iswift-go-project/internal/dashboard/dto"
	orderUsecase "iswift-go-project/internal/order/usecase"
	productUsecase "iswift-go-project/internal/product/usecase"
	userUsecase "iswift-go-project/internal/user/usecase"
)

type DashboardUseCase interface {
	GetDataDashboard() dto.DashboardResponseBody
}

type DashboardUseCaseImpl struct {
	userUseCase    userUsecase.UserUseCase
	adminUseCase   adminUsecase.AdminUseCase
	orderUseCase   orderUsecase.OrderUseCase
	productUseCase productUsecase.ProductUseCase
}

// GetDataDashboard implements DashboardUseCase.
func (usecase *DashboardUseCaseImpl) GetDataDashboard() dto.DashboardResponseBody {
	totalUser := usecase.userUseCase.Count()
	totalAdmin := usecase.adminUseCase.Count()
	totalOrder := usecase.orderUseCase.Count()
	totalProduct := usecase.productUseCase.Count()

	return dto.DashboardResponseBody{
		TotalUser:    int64(totalUser),
		TotalProduct: int64(totalProduct),
		TotalOrder:   int64(totalOrder),
		TotalAdmin:   int64(totalAdmin),
	}
}

func NewDashboardUseCase(
	userUseCase userUsecase.UserUseCase,
	adminUseCase adminUsecase.AdminUseCase,
	orderUseCase orderUsecase.OrderUseCase,
	productUseCase productUsecase.ProductUseCase,
) DashboardUseCase {
	return &DashboardUseCaseImpl{userUseCase, adminUseCase, orderUseCase, productUseCase}
}
