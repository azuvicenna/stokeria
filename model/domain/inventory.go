package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inventory struct {
	ID          string    `gorm:"type:char(36);primaryKey"`
	ProductID   string    `gorm:"type:char(36);not null;uniqueIndex:uk_product_warehouse"`
	WarehouseID string    `gorm:"type:char(36);not null;uniqueIndex:uk_product_warehouse"`
	Quantity    int       `gorm:"type:int;not null;default:0"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	// Relasi Belongs To
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
}

func (Inventory) TableName() string {
	return "inventory"
}

func (i *Inventory) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}
	return
}