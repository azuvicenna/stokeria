package web

import "time"

type InventoryResponse struct {
	ID        string             `json:"id"`
	Product   *ProductResponse   `json:"product,omitempty"`
	Warehouse *WarehouseResponse `json:"warehouse,omitempty"`
	Quantity  int                `json:"quantity"`
	UpdatedAt time.Time          `json:"updated_at"`
}