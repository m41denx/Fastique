package db

import (
	"gorm.io/gorm"
	"main/models"
)

type Role struct {
	gorm.Model
	Name       string
	Meta       models.JSONMap
	Privileges models.JSONMap
}
