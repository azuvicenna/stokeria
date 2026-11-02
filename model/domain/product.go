package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID         string    `gorm:"type:char(36);primaryKey"`
	SKU        string    `gorm:"column:sku;type:varchar(100);unique;not null"`
	CategoryID string    `gorm:"type:char(36);not null"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Price      float64   `gorm:"type:decimal(15,2);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	// Relasi Belongs To ke Category
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return
}
