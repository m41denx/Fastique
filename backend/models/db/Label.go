package db

import (
	"gorm.io/gorm"
	"main/models"
)

type Label struct {
	gorm.Model
	Name     string         `json:"name"`
	Template string         `json:"template"`
	Fields   models.JSONMap `json:"fields"`
}
