package product

import "time"

type ListProductRequest struct {
	Page     int    `json:"page" query:"page"`
	PerPage  int    `json:"per_page" query:"per_page"`
	Search   string `json:"search" query:"search"`
	IsActive *bool  `json:"is_active" query:"is_active"`
}

type ListProductResponse struct {
	ID             uint      `json:"id"`
	SKU            string    `json:"sku"`
	Name           string    `json:"name"`
	Barcode        *string   `json:"barcode"`
	Description    *string   `json:"description"`
	BaseUOM        string    `json:"base_uom"`
	Weight         *float64  `json:"weight"`
	Volume         *float64  `json:"volume"`
	IsBatchManaged *bool     `json:"is_batch_managed"`
	IsSerialized   *bool     `json:"is_serialized"`
	IsActive       *bool     `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	SKU            string   `json:"sku" validate:"required"`
	Name           string   `json:"name" validate:"required"`
	Barcode        *string  `json:"barcode"`
	Description    *string  `json:"description"`
	BaseUOM        string   `json:"base_uom" validate:"required"`
	Weight         *float64 `json:"weight" validate:"required"`
	Volume         *float64 `json:"volume" validate:"required"`
	IsBatchManaged *bool    `json:"is_batch_managed" validate:"required"`
	IsSerialized   *bool    `json:"is_serialized" validate:"required"`
	IsActive       *bool    `json:"is_active" validate:"required"`
}
