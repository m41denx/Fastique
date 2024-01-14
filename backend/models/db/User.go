package db

import (
	"gorm.io/gorm"
	"main/models"
)

type User struct {
	gorm.Model
	Username string         `json:"username"`
	PassHash string         `json:"-"`
	Meta     models.JSONMap `json:"meta"`
	Role     Role           `json:"role"`
	RoleID   uint           `json:"-"`
}
