package profile

import (
	userEntity "iswift-go-project/internal/user/entity"
)

type ProfilResponseBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func CreateProfileResponse(user userEntity.User) ProfilResponseBody {
	isVerified := false

	if user.EmailVerifiedAt.Valid {
		isVerified = true
	}

	return ProfilResponseBody{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: isVerified,
	}
}
