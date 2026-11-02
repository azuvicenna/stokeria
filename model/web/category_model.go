package web

import "time"

type CategoryCreateRequest struct {
	Name        string `validate:"required,max=150" json:"name"`
	Description string `json:"description"` // Optional, tidak ada tag required
}

type CategoryUpdateRequest struct {
	ID          string `validate:"required,uuid4" json:"id"`
	Name        string `validate:"required,max=150" json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}