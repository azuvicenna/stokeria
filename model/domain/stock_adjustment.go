package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockAdjustment struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	ProductID   string    `gorm:"type:char(36);not null"`
	WarehouseID string    `gorm:"type:char(36);not null"`
	UserID      string    `gorm:"type:char(36);not null"`
	OldQuantity int       `gorm:"type:int;not null"`
	NewQuantity int       `gorm:"type:int;not null"`
	Reason      string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	// Relasi Belongs To
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
}

func (s *StockAdjustment) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return
}
