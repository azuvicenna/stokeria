package web

import "time"

type WarehouseCreateRequest struct {
	Code    string `validate:"required,max=50" json:"code"`
	Name    string `validate:"required,max=255" json:"name"`
	Address string `json:"address"`
}

type WarehouseUpdateRequest struct {
	ID      string `validate:"required,uuid4" json:"id"`
	Code    string `validate:"required,max=50" json:"code"`
	Name    string `validate:"required,max=255" json:"name"`
	Address string `json:"address"`
}

type WarehouseResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}