package auth

import "time"

type UserRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type UserLogin struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	FullName  string     `json:"full_name"`
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	IsActive  *bool      `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Roles     []UserRole `json:"roles"`
}

type LoginResponse struct {
	User  UserLogin `json:"user"`
	Token string    `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
