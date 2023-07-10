package user

import (
	entity "iswift-go-project/internal/user/entity"
	repository "iswift-go-project/internal/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUsecase = UserUseCaseImpl{repository: userRepository}

func TestUserUsecase_FindByIdSuccess(t *testing.T) {
	userData := entity.User{
		ID:   1,
		Name: "test user",
	}

	userRepository.Mock.On("FindById", 1).Return(userData)

	user, err := userUsecase.FindById(1)

	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestUserUsecase_FindByIdNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", 2).Return(nil)

	user, err := userUsecase.FindById(2)

	assert.Nil(t, err)
	assert.Nil(t, user)
}
