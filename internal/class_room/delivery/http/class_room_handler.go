package class_room

import (
	usecase "iswift-go-project/internal/class_room/usecase"
	"iswift-go-project/internal/middleware"
	"iswift-go-project/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassRoomHandler struct {
	usecase usecase.ClassRoomUseCase
}

func (handler *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomHandler := r.Group("/api/v1")

	classRoomHandler.Use(middleware.AuthJwt)
	{
		classRoomHandler.GET("/class_rooms", handler.FindAllByUserId)
	}
}
func NewClassRoomHandler(usecase usecase.ClassRoomUseCase) *ClassRoomHandler {
	return &ClassRoomHandler{usecase}
}

func (handler *ClassRoomHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindAllByUserId(offset, limit, int(user.ID))

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
