package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name" validation:"required" gorm:"unique"`
	Permissions []Permission `json:"permissions" gorm:"many2many:roles_permissions"`
}
