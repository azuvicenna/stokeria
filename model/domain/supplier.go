package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID            string    `gorm:"type:char(36);primaryKey"`
	Name          string    `gorm:"type:varchar(255);not null"`
	ContactPerson string    `gorm:"type:varchar(150)"`
	Phone         string    `gorm:"type:varchar(50)"`
	Address       string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return
}
