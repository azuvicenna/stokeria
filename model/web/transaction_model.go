package web

import "time"

// Detail Transaksi (Digunakan oleh Inbound & Outbound)
type TransactionDetailCreateRequest struct {
	ProductID string  `validate:"required,uuid4" json:"product_id"`
	Quantity  int     `validate:"required,min=1" json:"quantity"`
	UnitPrice float64 `validate:"required,min=0" json:"unit_price"`
}

type TransactionDetailResponse struct {
	ID        string           `json:"id"`
	Product   *ProductResponse `json:"product,omitempty"`
	Quantity  int              `json:"quantity"`
	UnitPrice float64          `json:"unit_price"`
}

// Inbound Transaction
type InboundTransactionCreateRequest struct {
	SupplierID      string                           `validate:"required,uuid4" json:"supplier_id"`
	TransactionDate string                           `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"transaction_date"` // Standar RFC3339
	// Tag dive digunakan untuk memvalidasi isi di dalam array/slice
	Details         []TransactionDetailCreateRequest `validate:"required,min=1,dive" json:"details"` 
}

type InboundTransactionResponse struct {
	ID              string                      `json:"id"`
	ReferenceCode   string                      `json:"reference_code"`
	Supplier        *SupplierResponse           `json:"supplier,omitempty"`
	User            *UserResponse               `json:"user,omitempty"`
	TransactionDate time.Time                   `json:"transaction_date"`
	Status          string                      `json:"status"`
	Details         []TransactionDetailResponse `json:"details,omitempty"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}

// Outbound Transaction
type OutboundTransactionCreateRequest struct {
	Destination     string                           `validate:"required,max=255" json:"destination"`
	TransactionDate string                           `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"transaction_date"`
	Details         []TransactionDetailCreateRequest `validate:"required,min=1,dive" json:"details"`
}

type OutboundTransactionResponse struct {
	ID              string                      `json:"id"`
	ReferenceCode   string                      `json:"reference_code"`
	Destination     string                      `json:"destination"`
	User            *UserResponse               `json:"user,omitempty"`
	TransactionDate time.Time                   `json:"transaction_date"`
	Status          string                      `json:"status"`
	Details         []TransactionDetailResponse `json:"details,omitempty"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}