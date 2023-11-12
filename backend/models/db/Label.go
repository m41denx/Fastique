package db

import (
	"gorm.io/gorm"
	"main/models"
)

type Label struct {
	gorm.Model
	Name     string
	Template string
	Fields   models.JSONMap
}
