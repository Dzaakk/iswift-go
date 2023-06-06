package user

import (
	"errors"
	dto "iswift-go-project/internal/user/dto"
	entity "iswift-go-project/internal/user/entity"
	repository "iswift-go-project/internal/user/repository"
	utils "iswift-go-project/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	FindAll(offset, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(userDto dto.UserRequestBody) (*entity.User, error)
	Update(userDto dto.UserRequestBody) (*entity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

// FindByEmail implements UserUseCase
func (usecase *UserUseCaseImpl) FindByEmail(email string) (*entity.User, error) {
	return usecase.repository.FindByEmail(email)
}

// Create implements UserUseCase
func (usecase *UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {
	//Find by email

	checkUser, err := usecase.repository.FindByEmail(*userDto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, errors.New("email sudah pernah terdaftar")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	// Create data
	dataUser, err := usecase.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

// Delete implements UserUseCase
func (*UserUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements UserUseCase
func (*UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// FindById implements UserUseCase
func (usecase *UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	return usecase.repository.FindById(id)

}

// Update implements UserUseCase
func (*UserUseCaseImpl) Update(userDto dto.UserRequestBody) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
