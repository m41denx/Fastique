package db

import (
	"gorm.io/gorm"
	"main/models"
)

type User struct {
	gorm.Model
	Username string
	PassHash string
	Meta     models.JSONMap
	Role     Role
}
