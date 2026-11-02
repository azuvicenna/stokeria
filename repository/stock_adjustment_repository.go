package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type StockAdjustmentRepository interface {
	Create(ctx context.Context, adjustment *domain.StockAdjustment) error
	FindById(ctx context.Context, id string) (*domain.StockAdjustment, error)
	FindAll(ctx context.Context) ([]domain.StockAdjustment, error)
}

type stockAdjustmentRepositoryImpl struct {
	db            *gorm.DB
	inventoryRepo InventoryRepository
}

func NewStockAdjustmentRepository(db *gorm.DB, inventoryRepo InventoryRepository) StockAdjustmentRepository {
	return &stockAdjustmentRepositoryImpl{db: db, inventoryRepo: inventoryRepo}
}

func (r *stockAdjustmentRepositoryImpl) Create(ctx context.Context, adjustment *domain.StockAdjustment) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Simpan riwayat penyesuaian stok
		if err := tx.Create(adjustment).Error; err != nil {
			return err
		}

		// 2. Cari data inventory yang ada
		var inventory domain.Inventory
		err := tx.Where("product_id = ? AND warehouse_id = ?", adjustment.ProductID, adjustment.WarehouseID).First(&inventory).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				inventory = domain.Inventory{
					ProductID:   adjustment.ProductID,
					WarehouseID: adjustment.WarehouseID,
					Quantity:    adjustment.NewQuantity,
				}
				return tx.Create(&inventory).Error
			}
			return err
		}

		// 3. Update kuantitas stok inventory
		inventory.Quantity = adjustment.NewQuantity
		return tx.Save(&inventory).Error
	})
}

func (r *stockAdjustmentRepositoryImpl) FindById(ctx context.Context, id string) (*domain.StockAdjustment, error) {
	var adjustment domain.StockAdjustment
	err := r.db.WithContext(ctx).
		Preload("Product").
		Preload("Warehouse").
		Preload("User").
		First(&adjustment, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stock adjustment not found")
		}
		return nil, err
	}
	return &adjustment, nil
}

func (r *stockAdjustmentRepositoryImpl) FindAll(ctx context.Context) ([]domain.StockAdjustment, error) {
	var adjustments []domain.StockAdjustment
	err := r.db.WithContext(ctx).
		Preload("Product").
		Preload("Warehouse").
		Preload("User").
		Find(&adjustments).Error

	return adjustments, err
}
