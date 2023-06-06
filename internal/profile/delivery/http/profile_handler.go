package profile

import (
	middleware "iswift-go-project/internal/middleware"
	useCase "iswift-go-project/internal/profile/usecase"
	utils "iswift-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	useCase useCase.ProfileUseCase
}

func NewProfileHanlder(useCase useCase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{useCase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	authorized := r.Group("api/v1")

	authorized.Use(middleware.AuthJwt)
	{
		authorized.GET("/profiles", handler.GetProfile)
	}
}

func (handler *ProfileHandler) GetProfile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	// Get Profile
	profile, err := handler.useCase.GetProfile(int(user.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", profile))
}
