package web

import "time"

type UserCreateRequest struct {
	Email    string `validate:"required,email,max=255" json:"email"`
	Password string `validate:"required,min=6,max=255" json:"password"`
	Name     string `validate:"required,max=255" json:"name"`
	Role     string `validate:"required,oneof=admin warehouse_staff" json:"role"`
}

type UserUpdateRequest struct {
	ID    string `validate:"required,uuid4" json:"id"`
	Email string `validate:"required,email,max=255" json:"email"`
	Name  string `validate:"required,max=255" json:"name"`
	Role  string `validate:"required,oneof=admin warehouse_staff" json:"role"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Token     string       `json:"token"`
	TokenType string       `json:"token_type"`
	User      UserResponse `json:"user"`
}
