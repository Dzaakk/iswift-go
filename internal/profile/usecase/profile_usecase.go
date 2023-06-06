package profile

import (
	dto "iswift-go-project/internal/profile/dto"
	userUseCase "iswift-go-project/internal/user/usecase"
)

type ProfileUseCase interface {
	GetProfile(id int) (*dto.ProfilResponseBody, error)
}

type ProfileUseCaseImpl struct {
	userUseCase userUseCase.UserUseCase
}

// GetProfile implements ProfileUseCase.
func (usecase *ProfileUseCaseImpl) GetProfile(id int) (*dto.ProfilResponseBody, error) {
	// Get Profile
	user, err := usecase.userUseCase.FindById(id)

	if err != nil {
		return nil, err
	}

	UserResponse := dto.CreateProfileResponse(*user)

	return &UserResponse, nil
}

func NewProfileUseCase(userUseCase userUseCase.UserUseCase) ProfileUseCase {
	return &ProfileUseCaseImpl{userUseCase}
}
