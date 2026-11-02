package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	ID              string  `gorm:"type:char(36);primaryKey"`
	TransactionID   string  `gorm:"type:char(36);not null;index:idx_transaction_details_polymorphic"`
	TransactionType string  `gorm:"type:varchar(50);not null;index:idx_transaction_details_polymorphic"`
	ProductID       string  `gorm:"type:char(36);not null"`
	Quantity        int     `gorm:"type:int;not null"`
	UnitPrice       float64 `gorm:"type:decimal(15,2);not null"`

	// Relasi Belongs To
	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}

func (t *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return
}
