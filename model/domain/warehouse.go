package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Code      string    `gorm:"type:varchar(50);unique;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Address   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (w *Warehouse) BeforeCreate(tx *gorm.DB) (err error) {
	if w.ID == "" {
		w.ID = uuid.New().String()
	}
	return
}
