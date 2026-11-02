package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OutboundTransaction struct {
	ID              string    `gorm:"type:char(36);primaryKey"`
	ReferenceCode   string    `gorm:"type:varchar(100);unique;not null"`
	Destination     string    `gorm:"type:varchar(255);not null"`
	UserID          string    `gorm:"type:char(36);not null"`
	TransactionDate time.Time `gorm:"type:timestamp;not null"`
	Status          string    `gorm:"type:varchar(50);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	// Relasi Belongs To
	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (o *OutboundTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return
}