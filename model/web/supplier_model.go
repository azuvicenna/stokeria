package web

import "time"

type SupplierCreateRequest struct {
	Name          string `validate:"required,max=255" json:"name"`
	ContactPerson string `validate:"max=150" json:"contact_person"`
	Phone         string `validate:"max=50" json:"phone"`
	Address       string `json:"address"`
}

type SupplierUpdateRequest struct {
	ID            string `validate:"required,uuid4" json:"id"`
	Name          string `validate:"required,max=255" json:"name"`
	ContactPerson string `validate:"max=150" json:"contact_person"`
	Phone         string `validate:"max=50" json:"phone"`
	Address       string `json:"address"`
}

type SupplierResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	ContactPerson string    `json:"contact_person"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}