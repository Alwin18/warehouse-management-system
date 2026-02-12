package supplier

import "time"

type ListSupplierRequest struct {
	Page     int    `json:"page" query:"page"`
	PerPage  int    `json:"per_page" query:"per_page"`
	Search   string `json:"search" query:"search"`
	IsActive *bool  `json:"is_active" query:"is_active"`
}

type ListSupplierResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address   *string   `json:"address"`
	City      *string   `json:"city"`
	Country   *string   `json:"country"`
	Phone     *string   `json:"phone"`
	Email     *string   `json:"email"`
	IsActive  *bool     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSupplierRequest struct {
	Code     string  `json:"code" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Address  *string `json:"address"`
	City     *string `json:"city"`
	Country  *string `json:"country"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
	IsActive *bool   `json:"is_active"`
}
