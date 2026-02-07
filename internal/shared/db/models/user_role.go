package models

import "time"

type UserRole struct {
	UserID    uint      `gorm:"not null;primaryKey;index:idx_user_roles_user_id"`
	RoleID    uint      `gorm:"not null;primaryKey;index:idx_user_roles_role_id"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Role *Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
