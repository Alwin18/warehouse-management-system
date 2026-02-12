package models

import "time"

type ProductUOM struct {
	ID                uint      `gorm:"primary_key" json:"id"`
	ProductID         uint      `gorm:"not null;uniqueIndex:ux_product_uom;index:idx_prod_uom" json:"product_id"`
	UOMID             uint      `gorm:"not null;uniqueIndex:ux_product_uom;column:uom_id" json:"uom_id"`
	ConversionToBase  float64   `gorm:"type:decimal(18,6);not null;column:conversion_to_base" json:"conversion_to_base"`
	IsBaseUOM         *bool     `gorm:"type:boolean;not null;default:false;column:is_base_uom" json:"is_base_uom"`
	IsDefaultSales    *bool     `gorm:"type:boolean;not null;default:false;column:is_default_sales" json:"is_default_sales"`
	IsDefaultPurchase *bool     `gorm:"type:boolean;not null;default:false;column:is_default_purchase" json:"is_default_purchase"`
	CreatedAt         time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	Product *Product       `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product,omitempty"`
	UOM     *UnitOfMeasure `gorm:"foreignKey:UOMID;constraint:OnDelete:RESTRICT;" json:"uom,omitempty"`
}

func (ProductUOM) TableName() string {
	return "product_uoms"
}
