package web

import "time"

type ProductCreateRequest struct {
	SKU        string  `validate:"required,max=100" json:"sku"`
	CategoryID string  `validate:"required,uuid4" json:"category_id"`
	Name       string  `validate:"required,max=255" json:"name"`
	Price      float64 `validate:"required,min=0" json:"price"`
}

type ProductUpdateRequest struct {
	ID         string  `validate:"required,uuid4" json:"id"`
	SKU        string  `validate:"required,max=100" json:"sku"`
	CategoryID string  `validate:"required,uuid4" json:"category_id"`
	Name       string  `validate:"required,max=255" json:"name"`
	Price      float64 `validate:"required,min=0" json:"price"`
}

type ProductResponse struct {
	ID         string            `json:"id"`
	SKU        string            `json:"sku"`
	Name       string            `json:"name"`
	Price      float64           `json:"price"`
	Category   *CategoryResponse `json:"category,omitempty"` // Relasi disertakan dalam response
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}