//go:build wireinject
// +build wireinject

package user

import (
	handler "iswift-go-project/internal/user/delivery/http"
	repository "iswift-go-project/internal/user/repository"
	usecase "iswift-go-project/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		handler.NewUserHandler,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
	)

	return &handler.UserHandler{}
}
