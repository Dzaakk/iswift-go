package class_room

import (
	"database/sql"
	classRoomEntity "iswift-go-project/internal/class_room/entity"
	productEntity "iswift-go-project/internal/product/entity"
	userEntity "iswift-go-project/internal/user/entity"

	"gorm.io/gorm"
)

type ClassRoomResponseBody struct {
	ID        int64                  `json:"id"`
	User      *userEntity.User       `json:"user"`
	Product   *productEntity.Product `json:"product"`
	CreatedBy *userEntity.User       `json:"created_by" `
	UpdatedBy *userEntity.User       `json:"updated_by" `
	CreatedAt sql.NullTime           `json:"created_at"`
	UpdatedAt sql.NullTime           `json:"updated_at"`
	DeletedAt gorm.DeletedAt         `json:"deleted_at"`
}

func CreateClassRoomResponse(enitty classRoomEntity.ClassRoom) ClassRoomResponseBody {
	return ClassRoomResponseBody{
		ID:        enitty.ID,
		User:      enitty.User,
		Product:   enitty.Product,
		CreatedBy: enitty.CreatedBy,
		UpdatedBy: enitty.UpdatedBy,
		CreatedAt: enitty.CreatedAt,
		UpdatedAt: enitty.UpdatedAt,
		DeletedAt: enitty.DeletedAt,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(entity []classRoomEntity.ClassRoom) ClassRoomListResponse {
	classRoomResp := ClassRoomListResponse{}

	for _, classRoom := range entity {
		classRoom.Product.VideoURL = classRoom.Product.Video

		classRoomResponse := CreateClassRoomResponse(classRoom)
		classRoomResp = append(classRoomResp, classRoomResponse)
	}

	return classRoomResp
}
