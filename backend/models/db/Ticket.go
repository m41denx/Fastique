package db

import (
	"gorm.io/gorm"
	"main/models"
	"time"
)

type Ticket struct {
	gorm.Model
	Name      string
	Branch    Branch
	Label     Label
	BeginTime time.Time
	EndTime   time.Time
	Servant   User
	meta      models.JSONMap
}
