//go:build wireinject
// +build wireinject

package oauth

import (
	adminRepository "iswift-go-project/internal/admin/repository"
	adminUseCase "iswift-go-project/internal/admin/usecase"
	oauthHandler "iswift-go-project/internal/oauth/delivery/http"
	oauthRepository "iswift-go-project/internal/oauth/repository"
	oauthUseCase "iswift-go-project/internal/oauth/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUseCase "iswift-go-project/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		oauthUseCase.NewOauthUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		adminRepository.NewAdminRepository,
		adminUseCase.NewAdminUseCase,
	)

	return &oauthHandler.OauthHandler{}
}
