package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type InboundTransactionRepository interface {
	Create(ctx context.Context, transaction *domain.InboundTransaction, details []domain.TransactionDetail) error
	FindById(ctx context.Context, id string) (*domain.InboundTransaction, error)
	FindAll(ctx context.Context) ([]domain.InboundTransaction, error)
}

type inboundTransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewInboundTransactionRepository(db *gorm.DB) InboundTransactionRepository {
	return &inboundTransactionRepositoryImpl{db: db}
}

func (r *inboundTransactionRepositoryImpl) Create(ctx context.Context, transaction *domain.InboundTransaction, details []domain.TransactionDetail) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Simpan Header Transaksi
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// 2. Assign TransactionID dan TransactionType ke masing-masing Detail
		for i := range details {
			details[i].TransactionID = transaction.ID
			details[i].TransactionType = "INBOUND"
		}

		// 3. Simpan Details secara Bulk Insert
		if err := tx.Create(&details).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *inboundTransactionRepositoryImpl) FindById(ctx context.Context, id string) (*domain.InboundTransaction, error) {
	var transaction domain.InboundTransaction
	err := r.db.WithContext(ctx).
		Preload("Supplier").
		Preload("User").
		First(&transaction, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inbound transaction not found")
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *inboundTransactionRepositoryImpl) FindAll(ctx context.Context) ([]domain.InboundTransaction, error) {
	var transactions []domain.InboundTransaction
	err := r.db.WithContext(ctx).
		Preload("Supplier").
		Preload("User").
		Find(&transactions).Error

	return transactions, err
}