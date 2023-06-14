package order

import (
	entity "iswift-go-project/internal/order/entity"
	"iswift-go-project/pkg/utils"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(offset, limit int) []entity.Order
	FindAllByUserId(offset, limit, userId int) []entity.Order
	FindOneByExternalId(externalId string) (*entity.Order, error)
	FindById(id int) (*entity.Order, error)
	Create(entity entity.Order) (*entity.Order, error)
	Update(entity entity.Order) (*entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

// FindAllByUserId implements OrderRepository.
func (repository *OrderRepositoryImpl) FindAllByUserId(offset int, limit int, userId int) []entity.Order {
	var orders []entity.Order

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("OrderDetails.Product").
		Where("user_id = ?", userId).
		Find(&orders)

	return orders
}

// FindOneByExternalId implements OrderRepository.
func (repository *OrderRepositoryImpl) FindOneByExternalId(externalId string) (*entity.Order, error) {
	var order entity.Order

	// menampilkan data order dengan child order details
	// 1 order dapat mempunyai banyak order details
	// one to many
	if err := repository.db.
		Preload("OrderDetails.Product").
		Where("external_id = ?", externalId).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

// Update implements OrderRepository.
func (repository *OrderRepositoryImpl) Update(entity entity.Order) (*entity.Order, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Create implements OrderRepository.
func (repository *OrderRepositoryImpl) Create(entity entity.Order) (*entity.Order, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAll implements OrderRepository.
func (repository *OrderRepositoryImpl) FindAll(offset int, limit int) []entity.Order {
	var orders []entity.Order

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&orders)

	return orders
}

// FindById implements OrderRepository.
func (repository *OrderRepositoryImpl) FindById(id int) (*entity.Order, error) {
	var order entity.Order

	if err := repository.db.First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db}
}
