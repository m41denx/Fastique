package db

import (
	"gorm.io/gorm"
	"main/models"
)

type Role struct {
	gorm.Model
	Name       string         `json:"name"`
	Meta       models.JSONMap `json:"meta"`
	Privileges models.JSONMap `json:"privileges"`
}
