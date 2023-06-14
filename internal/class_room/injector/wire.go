//go:build wireinject

// +wireinject

package class_room

import (
	handler "iswift-go-project/internal/class_room/delivery/http"
	repository "iswift-go-project/internal/class_room/repository"
	usecase "iswift-go-project/internal/class_room/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ClassRoomHandler {
	wire.Build(
		handler.NewClassRoomHandler,
		repository.NewClassRoomRepository,
		usecase.NewClassRoomUseCase,
	)

	return &handler.ClassRoomHandler{}
}
