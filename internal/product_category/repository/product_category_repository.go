package product_category

import (
	entity "iswift-go-project/internal/product_category/entity"
	"iswift-go-project/pkg/utils"

	"gorm.io/gorm"
)

type ProductcategoryRepository interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindById(id int) (*entity.ProductCategory, error)
	Create(entity entity.ProductCategory) (*entity.ProductCategory, error)
	Update(entity entity.ProductCategory) (*entity.ProductCategory, error)
	Delete(entity entity.ProductCategory) error
}

type ProductcategoryRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ProductcategoryRepository.
func (repository *ProductcategoryRepositoryImpl) Create(entity entity.ProductCategory) (*entity.ProductCategory, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements ProductcategoryRepository.
func (repository *ProductcategoryRepositoryImpl) Delete(entity entity.ProductCategory) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductcategoryRepository.
func (repository *ProductcategoryRepositoryImpl) FindAll(offset int, limit int) []entity.ProductCategory {
	var productCategories []entity.ProductCategory

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories)

	return productCategories
}

// FindById implements ProductcategoryRepository.
func (repository *ProductcategoryRepositoryImpl) FindById(id int) (*entity.ProductCategory, error) {
	var productCategory entity.ProductCategory

	if err := repository.db.First(&productCategory, id).Error; err != nil {
		return nil, err
	}

	return &productCategory, nil
}

// Update implements ProductcategoryRepository.
func (repository *ProductcategoryRepositoryImpl) Update(entity entity.ProductCategory) (*entity.ProductCategory, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductcategoryRepository {
	return &ProductcategoryRepositoryImpl{db}
}
