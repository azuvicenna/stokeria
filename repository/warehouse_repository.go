package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse *domain.Warehouse) error
	Update(ctx context.Context, warehouse *domain.Warehouse) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*domain.Warehouse, error)
	FindByCode(ctx context.Context, code string) (*domain.Warehouse, error)
	FindAll(ctx context.Context) ([]domain.Warehouse, error)
}

type warehouseRepositoryImpl struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepositoryImpl{db: db}
}

func (r *warehouseRepositoryImpl) Create(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *warehouseRepositoryImpl) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Save(warehouse).Error
}

func (r *warehouseRepositoryImpl) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&domain.Warehouse{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("warehouse not found")
	}
	return nil
}

func (r *warehouseRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("warehouse not found")
		}
		return nil, err
	}
	return &warehouse, nil
}

func (r *warehouseRepositoryImpl) FindByCode(ctx context.Context, code string) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, "code = ?", code).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("warehouse not found")
		}
		return nil, err
	}
	return &warehouse, nil
}

func (r *warehouseRepositoryImpl) FindAll(ctx context.Context) ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	err := r.db.WithContext(ctx).Find(&warehouses).Error
	return warehouses, err
}