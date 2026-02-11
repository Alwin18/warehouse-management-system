package models

import "time"

type UnitOfMeasure struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Code        string    `gorm:"type:varchar(20);not null;uniqueIndex:ux_uom_code" json:"code"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Symbol      *string   `gorm:"type:varchar(10)" json:"symbol"`
	Category    string    `gorm:"type:varchar(50);not null" json:"category"` // WEIGHT, VOLUME, COUNT, LENGTH, AREA
	Description *string   `gorm:"type:text" json:"description"`
	IsActive    *bool     `gorm:"type:boolean;not null;default:true;column:is_active" json:"is_active"`
	CreatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *UnitOfMeasure) TableName() string {
	return "unit_of_measures"
}
