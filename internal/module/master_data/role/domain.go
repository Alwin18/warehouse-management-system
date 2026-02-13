package role

import "time"

type ListRoleRequest struct {
	Page    int    `json:"page" query:"page"`
	PerPage int    `json:"per_page" query:"per_page"`
	Search  string `json:"search" query:"search"`
}

type ListRoleResponse struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateRoleRequest struct {
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}
