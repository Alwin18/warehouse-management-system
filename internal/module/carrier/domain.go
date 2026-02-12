package carrier

import "time"

type ListCarrierRequest struct {
	Page     int    `json:"page" query:"page"`
	PerPage  int    `json:"per_page" query:"per_page"`
	Search   string `json:"search" query:"search"`
	IsActive *bool  `json:"is_active" query:"is_active"`
}

type ListCarrierResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	IsActive  *bool     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCarrierRequest struct {
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required"`
	IsActive *bool  `json:"is_active"`
}
