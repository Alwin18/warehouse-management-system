package purchaseorder

import "time"

type ListPurchaseOrderRequest struct {
	Page     int    `query:"page"`
	PerPage  int    `query:"per_page"`
	Search   string `query:"search"`
	IsActive *bool  `query:"is_active"`
}

type ListPurchaseOrderResponse struct {
	ID           uint       `json:"id"`
	PONumber     string     `json:"po_number"`
	SupplierID   uint       `json:"supplier_id"`
	WarehouseID  uint       `json:"warehouse_id"`
	Status       string     `json:"status"`
	OrderDate    time.Time  `json:"order_date"`
	ExpectedDate *time.Time `json:"expected_date"`
	Currency     *string    `json:"currency"`
	TotalAmount  *float64   `json:"total_amount"`
	CreatedBy    string     `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CreatePurchaseOrderRequest struct {
	SupplierID   uint       `json:"supplier_id" validate:"required"`
	WarehouseID  uint       `json:"warehouse_id" validate:"required"`
	Status       string     `json:"status" validate:"required"`
	OrderDate    time.Time  `json:"order_date" validate:"required"`
	ExpectedDate *time.Time `json:"expected_date"`
	Currency     *string    `json:"currency"`
	TotalAmount  *float64   `json:"total_amount"`
}
