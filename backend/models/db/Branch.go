package db

import (
	"main/models"
)

type Branch struct {
	ID            uint
	Name          string         `json:"name"`
	ServicePoints []ServicePoint `json:"service_points"`
	Meta          models.JSONMap `json:"meta"`
}
