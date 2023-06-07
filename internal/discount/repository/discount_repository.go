package discount

import (
	entity "iswift-go-project/internal/discount/entity"
	"iswift-go-project/pkg/utils"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(offset, limit int) []entity.Discount
	FindByID(id int) (*entity.Discount, error)
	FindByCode(code string) (*entity.Discount, error)
	Create(entity entity.Discount) (*entity.Discount, error)
	Update(entity entity.Discount) (*entity.Discount, error)
	Delete(entity entity.Discount) error
}

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

// Delete implements DiscountRepository.
func (repository *DiscountRepositoryImpl) Delete(entity entity.Discount) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// Create implements DiscountRepository.
func (repository *DiscountRepositoryImpl) Create(entity entity.Discount) (*entity.Discount, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAll implements DiscountRepository.
func (repository *DiscountRepositoryImpl) FindAll(offset int, limit int) []entity.Discount {
	var discounts []entity.Discount

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts)

	return discounts
}

// FindByCode implements DiscountRepository.
func (repository *DiscountRepositoryImpl) FindByCode(code string) (*entity.Discount, error) {
	var discount entity.Discount

	if err := repository.db.Where("code = ?", code).First(&discount).Error; err != nil {
		return nil, err
	}

	return &discount, nil
}

// FindByID implements DiscountRepository.
func (repository *DiscountRepositoryImpl) FindByID(id int) (*entity.Discount, error) {
	var discount entity.Discount

	if err := repository.db.First(&discount, id).Error; err != nil {
		return nil, err
	}

	return &discount, nil
}

// Update implements DiscountRepository.
func (repository *DiscountRepositoryImpl) Update(entity entity.Discount) (*entity.Discount, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db}
}
