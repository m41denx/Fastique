package db

import (
	"gorm.io/gorm"
	"main/models"
	"time"
)

type Ticket struct {
	gorm.Model
	Name           string         `json:"name"`
	ServicePoint   ServicePoint   `json:"sp"`
	ServicePointID uint           `json:"-"`
	Label          Label          `json:"label"`
	LabelID        uint           `json:"-"`
	BeginTime      time.Time      `json:"beginTime"`
	EndTime        time.Time      `json:"endTime"`
	Servant        User           `json:"servant"`
	ServantID      uint           `json:"-"`
	Meta           models.JSONMap `json:"meta"`
}
