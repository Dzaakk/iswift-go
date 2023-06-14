package webhook

import (
	dto "iswift-go-project/internal/webhook/dto"
	usecase "iswift-go-project/internal/webhook/usecase"
	"iswift-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	usecase usecase.WebhookUseCase
}

func NewWebHookHandler(usecase usecase.WebhookUseCase) *WebhookHandler {
	return &WebhookHandler{usecase}
}
func (handler *WebhookHandler) Route(r *gin.RouterGroup) {
	webhookHandler := r.Group("/api/v1")

	webhookHandler.POST("/webhooks", handler.Xendit)
}
func (handler *WebhookHandler) Xendit(ctx *gin.Context) {
	var input dto.WebhookRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err := handler.usecase.UpdatePayment(input.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))
}
