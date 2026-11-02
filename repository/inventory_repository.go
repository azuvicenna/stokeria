package repository

import (
	"context"
	"errors"
	"fmt"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(ctx context.Context, inventory *domain.Inventory) error
	Update(ctx context.Context, inventory *domain.Inventory) error
	FindById(ctx context.Context, id string) (*domain.Inventory, error)
	FindByProductAndWarehouse(ctx context.Context, productID string, warehouseID string) (*domain.Inventory, error)
	FindAll(ctx context.Context) ([]domain.Inventory, error)
	MutateStock(ctx context.Context, tx *gorm.DB, productID string, warehouseID string, quantityDelta int) (*domain.Inventory, error)
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

// MutateStock menambah (quantityDelta > 0) atau mengurangi (quantityDelta < 0) stok inventory secara atomik
func (r *inventoryRepositoryImpl) MutateStock(ctx context.Context, tx *gorm.DB, productID string, warehouseID string, quantityDelta int) (*domain.Inventory, error) {
	runner := r.db
	if tx != nil {
		runner = tx
	}

	var inventory domain.Inventory
	err := runner.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		First(&inventory).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if quantityDelta < 0 {
				return nil, fmt.Errorf("insufficient stock: product %s not found in warehouse %s", productID, warehouseID)
			}
			inventory = domain.Inventory{
				ProductID:   productID,
				WarehouseID: warehouseID,
				Quantity:    quantityDelta,
			}
			if err := runner.WithContext(ctx).Create(&inventory).Error; err != nil {
				return nil, err
			}
			return &inventory, nil
		}
		return nil, err
	}

	if inventory.Quantity+quantityDelta < 0 {
		return nil, fmt.Errorf("insufficient stock for product %s in warehouse %s: available %d, required %d", productID, warehouseID, inventory.Quantity, -quantityDelta)
	}

	inventory.Quantity += quantityDelta
	if err := runner.WithContext(ctx).Save(&inventory).Error; err != nil {
		return nil, err
	}

	return &inventory, nil
}
