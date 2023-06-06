//go:build wireinject
// +build wireinject

package admin

import (
	handler "iswift-go-project/internal/admin/delivery/http"
	repository "iswift-go-project/internal/admin/repository"
	usecase "iswift-go-project/internal/admin/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
	)
	return &handler.AdminHandler{}
}
