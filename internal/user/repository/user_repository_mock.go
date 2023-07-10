package user

import (
	"github.com/stretchr/testify/mock"

	entity "iswift-go-project/internal/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

// Count implements UserRepository.
func (repository *UserRepositoryMock) Count() int {
	panic("")
}

// FindByEmail implements UserRepository
func (repository *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	panic("")
}

// FindAll implements UserRepository
func (repository *UserRepositoryMock) FindAll(offset int, limit int) []entity.User {
	panic("")
}

// FindById implements UserRepository
func (repository *UserRepositoryMock) FindById(id int) (*entity.User, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, nil
	}

	user := arguments.Get(0).(entity.User)
	return &user, nil
}

// Create implements UserRepository
func (repository *UserRepositoryMock) Create(entity entity.User) (*entity.User, error) {
	panic("")
}

// Update implements UserRepository
func (repository *UserRepositoryMock) Update(entity entity.User) (*entity.User, error) {
	panic("")
}

// Delete implements UserRepository
func (repository *UserRepositoryMock) Delete(entity entity.User) error {
	panic("")
}
