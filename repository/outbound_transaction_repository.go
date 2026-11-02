package repository

import (
	"context"
	"errors"
	"stokeria/model/domain"

	"gorm.io/gorm"
)

type OutboundTransactionRepository interface {
	Create(ctx context.Context, transaction *domain.OutboundTransaction, details []domain.TransactionDetail) error
	FindById(ctx context.Context, id string) (*domain.OutboundTransaction, error)
	FindAll(ctx context.Context) ([]domain.OutboundTransaction, error)
}

type outboundTransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewOutboundTransactionRepository(db *gorm.DB) OutboundTransactionRepository {
	return &outboundTransactionRepositoryImpl{db: db}
}

func (r *outboundTransactionRepositoryImpl) Create(ctx context.Context, transaction *domain.OutboundTransaction, details []domain.TransactionDetail) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		for i := range details {
			details[i].TransactionID = transaction.ID
			details[i].TransactionType = "OUTBOUND"
		}

		if err := tx.Create(&details).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *outboundTransactionRepositoryImpl) FindById(ctx context.Context, id string) (*domain.OutboundTransaction, error) {
	var transaction domain.OutboundTransaction
	err := r.db.WithContext(ctx).
		Preload("User").
		First(&transaction, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("outbound transaction not found")
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *outboundTransactionRepositoryImpl) FindAll(ctx context.Context) ([]domain.OutboundTransaction, error) {
	var transactions []domain.OutboundTransaction
	err := r.db.WithContext(ctx).
		Preload("User").
		Find(&transactions).Error

	return transactions, err
}