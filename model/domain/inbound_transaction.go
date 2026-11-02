package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InboundTransaction struct {
	ID              string    `gorm:"type:char(36);primaryKey"`
	ReferenceCode   string    `gorm:"type:varchar(100);unique;not null"`
	SupplierID      string    `gorm:"type:char(36);not null"`
	WarehouseID     string    `gorm:"type:char(36);not null"`
	UserID          string    `gorm:"type:char(36);not null"`
	TransactionDate time.Time `gorm:"type:timestamp;not null"`
	Status          string    `gorm:"type:varchar(50);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	// Relasi Belongs To
	Supplier  Supplier  `gorm:"foreignKey:SupplierID;references:ID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
}

func (i *InboundTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}
	return
}
