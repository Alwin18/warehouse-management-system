package models

import "time"

type CustomerReturnLine struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	CustomerReturnID uint      `gorm:"not null;index:idx_crlines_cr" json:"customer_return_id"`
	SalesOrderLineID uint      `gorm:"not null;index:idx_crlines_sol" json:"sales_order_line_id"`
	ProductID        uint      `gorm:"not null;index:idx_crlines_product" json:"product_id"`
	UOMID            uint      `gorm:"not null;column:uom_id" json:"uom_id"`
	ReturnedQty      float64   `gorm:"type:decimal(18,3);not null;column:returned_qty" json:"returned_qty"`
	BatchID          *uint     `gorm:"index:idx_crlines_batch" json:"batch_id"`
	QCStatus         *string   `gorm:"type:varchar(50);column:qc_status" json:"qc_status"`
	ReturnReasonCode *string   `gorm:"type:varchar(50);column:return_reason_code" json:"return_reason_code"`
	CreatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	CustomerReturn *CustomerReturn `gorm:"foreignKey:CustomerReturnID;constraint:OnDelete:CASCADE;" json:"customer_return,omitempty"`
	SalesOrderLine *SalesOrderLine `gorm:"foreignKey:SalesOrderLineID;constraint:OnDelete:RESTRICT;" json:"sales_order_line,omitempty"`
	Product        *Product        `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT;" json:"product,omitempty"`
	UOM            *UnitOfMeasure  `gorm:"foreignKey:UOMID;constraint:OnDelete:RESTRICT;" json:"uom,omitempty"`
	Batch          *ProductBatch   `gorm:"foreignKey:BatchID;constraint:OnDelete:SET NULL;" json:"batch,omitempty"`
}

func (CustomerReturnLine) TableName() string {
	return "customer_return_lines"
}
