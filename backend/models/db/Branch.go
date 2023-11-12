package db

import (
	"gorm.io/gorm"
	"main/models"
)

type Branch struct {
	gorm.Model
	Name string
	Meta models.JSONMap
}
