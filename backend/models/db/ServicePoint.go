package db

import (
	"gorm.io/gorm"
	"main/models"
)

type ServicePoint struct {
	gorm.Model
	Branch            Branch
	Labels            []Label
	Available         bool
	User              User
	AvailabilityTimes models.JSONMap
}
