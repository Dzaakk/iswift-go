package dashboard

import (
	usecase "iswift-go-project/internal/dashboard/usecase"
	"iswift-go-project/internal/middleware"
	"iswift-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	usecase usecase.DashboardUseCase
}

func NewDashboardHandler(usecase usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{usecase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardHandler := r.Group("/api/v1")

	dashboardHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardHandler.GET("/dashboards", handler.GetDataDashboard)
	}
}
func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.usecase.GetDataDashboard()

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
