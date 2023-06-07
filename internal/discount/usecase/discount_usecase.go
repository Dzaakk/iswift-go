package discount

import (
	"database/sql"
	"errors"
	dto "iswift-go-project/internal/discount/dto"
	entity "iswift-go-project/internal/discount/entity"
	repository "iswift-go-project/internal/discount/repository"
)

type DiscountUseCase interface {
	FindAll(offset, limit int) []entity.Discount
	FindById(id int) (*entity.Discount, error)
	FindByCode(code string) (*entity.Discount, error)
	Create(dto dto.DiscountRequestBody) (*entity.Discount, error)
	Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, error)
	Delete(id int) error
	UpdateRemainingQuantity(id, quantity int, operator string) (*entity.Discount, error)
}

type DiscountUseCaseImpl struct {
	repository repository.DiscountRepository
}

// Create implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) Create(dto dto.DiscountRequestBody) (*entity.Discount, error) {
	discount := entity.Discount{
		Name:              dto.Name,
		Code:              dto.Code,
		Quantity:          dto.Quantity,
		RemainingQuantity: 0,
		Type:              dto.Type,
		Value:             dto.Value,
		StartDate: sql.NullTime{
			Time:  dto.StartDate,
			Valid: true,
		},
		EndDate: sql.NullTime{
			Time:  dto.EndDate,
			Valid: true,
		},
		CreatedByID: &dto.CreatedBy,
	}
	data, err := usecase.repository.Create(discount)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) Delete(id int) error {
	// Cari data discount berdasarkan id
	discount, err := usecase.repository.FindByID(id)

	if err != nil {
		return err
	}

	// Panggil repository delete
	err = usecase.repository.Delete(*discount)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) FindAll(offset int, limit int) []entity.Discount {
	return usecase.repository.FindAll(offset, limit)
}

// FindByCode implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) FindByCode(code string) (*entity.Discount, error) {
	return usecase.repository.FindByCode(code)
}

// FindById implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) FindById(id int) (*entity.Discount, error) {
	return usecase.repository.FindByID(id)
}

// Update implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, error) {
	// Cari data discount berdasarkan id
	discount, err := usecase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	discount.Name = dto.Name
	discount.Code = dto.Code
	discount.Quantity = dto.Quantity
	discount.RemainingQuantity = dto.RemainingQuantity
	discount.Type = dto.Type
	discount.Value = dto.Value
	discount.UpdatedByID = &dto.UpdatedBy
	discount.StartDate.Time = dto.StartDate
	discount.EndDate.Time = dto.EndDate

	updateDiscount, err := usecase.repository.Update(*discount)

	if err != nil {
		return nil, err
	}

	return updateDiscount, nil

}

// UpdateRemainingQuantity implements DiscountUseCase.
func (usecase *DiscountUseCaseImpl) UpdateRemainingQuantity(id int, quantity int, operator string) (*entity.Discount, error) {
	// Cari data discount berdasarkan id
	discount, err := usecase.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	if operator == "+" {
		discount.RemainingQuantity = discount.RemainingQuantity + int64(quantity)
	} else if operator == "-" {
		discount.RemainingQuantity = discount.RemainingQuantity - int64(quantity)
	} else {
		return nil, errors.New("operator belum dihandle")
	}

	updateDiscount, err := usecase.repository.Update(*discount)

	if err != nil {
		return nil, err
	}

	return updateDiscount, nil
}

func NewDiscountUseCase(repository repository.DiscountRepository) DiscountUseCase {
	return &DiscountUseCaseImpl{repository}
}
