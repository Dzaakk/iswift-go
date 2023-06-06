//go:build wireinject
// +build wireinject

package register

import (
	handler "iswift-go-project/internal/register/delivery/http"
	useCase "iswift-go-project/internal/register/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUseCase "iswift-go-project/internal/user/usecase"
	mail "iswift-go-project/pkg/mail/sendgrid"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		useCase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.RegisterHandler{}
}
