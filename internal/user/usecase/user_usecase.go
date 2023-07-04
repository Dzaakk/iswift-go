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
	Create(dto dto.UserRequestBody) (*entity.User, error)
	Update(id int, dto dto.UserRequestBody) (*entity.User, error)
	Count() int
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

// Count implements UserUseCase.
func (usecase *UserUseCaseImpl) Count() int {
	return usecase.repository.Count()
}

// FindByEmail implements UserUseCase
func (usecase *UserUseCaseImpl) FindByEmail(email string) (*entity.User, error) {
	return usecase.repository.FindByEmail(email)
}

// Create implements UserUseCase
func (usecase *UserUseCaseImpl) Create(dto dto.UserRequestBody) (*entity.User, error) {
	//Find by email

	checkUser, err := usecase.repository.FindByEmail(*dto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, errors.New("email sudah pernah terdaftar")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Name:         *dto.Name,
		Email:        *dto.Email,
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
func (usecase *UserUseCaseImpl) Delete(id int) error {
	user, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*user)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements UserUseCase
func (usecase *UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements UserUseCase
func (usecase *UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	return usecase.repository.FindById(id)

}

// Update implements UserUseCase
func (usecase *UserUseCaseImpl) Update(id int, dto dto.UserRequestBody) (*entity.User, error) {
	user, err := usecase.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		user.Name = *dto.Name
	}

	if dto.Email != nil {
		if user.Email != *dto.Email {
			user.Email = *dto.Email
		}
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
	}

	updateUser, err := usecase.repository.Update(*user)

	if err != nil {
		return nil, err
	}

	return updateUser, nil

}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
