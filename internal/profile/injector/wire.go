//go:build wireinject
// +build wireinject

package profile

import (
	handler "iswift-go-project/internal/profile/delivery/http"
	usecase "iswift-go-project/internal/profile/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUsecase "iswift-go-project/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProfileHandler {
	wire.Build(
		handler.NewProfileHanlder,
		usecase.NewProfileUseCase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
	)

	return &handler.ProfileHandler{}
}
