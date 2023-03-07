package entity

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string `json:"name" validation:"required" gorm:"unique"`
	Roles []Role `json:"roles" gorm:"many2many:roles_permissions"`
}
