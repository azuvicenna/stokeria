package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(ctx context.Context, supplier *domain.Supplier) error
	Update(ctx context.Context, supplier *domain.Supplier) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*domain.Supplier, error)
	FindAll(ctx context.Context) ([]domain.Supplier, error)
}

type supplierRepositoryImpl struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepositoryImpl{db: db}
}

func (r *supplierRepositoryImpl) Create(ctx context.Context, supplier *domain.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

func (r *supplierRepositoryImpl) Update(ctx context.Context, supplier *domain.Supplier) error {
	return r.db.WithContext(ctx).Save(supplier).Error
}

func (r *supplierRepositoryImpl) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&domain.Supplier{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("supplier not found")
	}
	return nil
}

func (r *supplierRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Supplier, error) {
	var supplier domain.Supplier
	err := r.db.WithContext(ctx).First(&supplier, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("supplier not found")
		}
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierRepositoryImpl) FindAll(ctx context.Context) ([]domain.Supplier, error) {
	var suppliers []domain.Supplier
	err := r.db.WithContext(ctx).Find(&suppliers).Error
	return suppliers, err
}