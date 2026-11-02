package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(ctx context.Context, inventory *domain.Inventory) error
	Update(ctx context.Context, inventory *domain.Inventory) error
	FindById(ctx context.Context, id string) (*domain.Inventory, error)
	FindByProductAndWarehouse(ctx context.Context, productID string, warehouseID string) (*domain.Inventory, error)
	FindAll(ctx context.Context) ([]domain.Inventory, error)
}

type inventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepositoryImpl{db: db}
}

func (r *inventoryRepositoryImpl) Create(ctx context.Context, inventory *domain.Inventory) error {
	return r.db.WithContext(ctx).Create(inventory).Error
}

func (r *inventoryRepositoryImpl) Update(ctx context.Context, inventory *domain.Inventory) error {
	return r.db.WithContext(ctx).Save(inventory).Error
}

func (r *inventoryRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Inventory, error) {
	var inventory domain.Inventory
	err := r.db.WithContext(ctx).
		Preload("Product").
		Preload("Warehouse").
		First(&inventory, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepositoryImpl) FindByProductAndWarehouse(ctx context.Context, productID string, warehouseID string) (*domain.Inventory, error) {
	var inventory domain.Inventory
	err := r.db.WithContext(ctx).
		Preload("Product").
		Preload("Warehouse").
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		First(&inventory).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found for given product and warehouse")
		}
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	err := r.db.WithContext(ctx).
		Preload("Product").
		Preload("Warehouse").
		Find(&inventories).Error

	return inventories, err
}