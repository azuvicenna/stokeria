package web

import "time"

type StockAdjustmentCreateRequest struct {
	ProductID   string `validate:"required,uuid4" json:"product_id"`
	WarehouseID string `validate:"required,uuid4" json:"warehouse_id"`
	OldQuantity int    `validate:"required,min=0" json:"old_quantity"`
	NewQuantity int    `validate:"required,min=0" json:"new_quantity"`
	Reason      string `validate:"required" json:"reason"`
}

type StockAdjustmentResponse struct {
	ID          string             `json:"id"`
	Product     *ProductResponse   `json:"product,omitempty"`
	Warehouse   *WarehouseResponse `json:"warehouse,omitempty"`
	User        *UserResponse      `json:"user,omitempty"`
	OldQuantity int                `json:"old_quantity"`
	NewQuantity int                `json:"new_quantity"`
	Reason      string             `json:"reason"`
	CreatedAt   time.Time          `json:"created_at"`
}