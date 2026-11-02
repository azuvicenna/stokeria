package repository

import (
	"context"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	// CreateMany digunakan untuk menyimpan banyak detail sekaligus (bulk insert)
	CreateMany(ctx context.Context, tx *gorm.DB, details []domain.TransactionDetail) error
	
	// FindByTransaction mencari detail berdasarkan ID transaksi dan tipenya ("INBOUND" / "OUTBOUND")
	FindByTransaction(ctx context.Context, transactionID string, transactionType string) ([]domain.TransactionDetail, error)
}

type transactionDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepositoryImpl{db: db}
}

// Perhatikan kita menerima *gorm.DB dari luar (parameter tx)
// Ini sangat penting agar CreateMany bisa berjalan di dalam siklus Database Transaction
// yang sama dengan pembuatan Inbound/Outbound Transaction utama.
func (r *transactionDetailRepositoryImpl) CreateMany(ctx context.Context, tx *gorm.DB, details []domain.TransactionDetail) error {
	return tx.WithContext(ctx).Create(&details).Error
}

func (r *transactionDetailRepositoryImpl) FindByTransaction(ctx context.Context, transactionID string, transactionType string) ([]domain.TransactionDetail, error) {
	var details []domain.TransactionDetail
	err := r.db.WithContext(ctx).
		Preload("Product").
		Where("transaction_id = ? AND transaction_type = ?", transactionID, transactionType).
		Find(&details).Error

	return details, err
}