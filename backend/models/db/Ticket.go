package db

import (
	"gorm.io/gorm"
	"main/models"
	"time"
)

type Ticket struct {
	gorm.Model
	Name           string         `json:"name"`
	Branch         Branch         `json:"branch"`
	BranchID       uint           `json:"-"`
	ServicePoint   ServicePoint   `json:"sp"` // Assign only after start
	ServicePointID uint           `json:"spid"`
	Label          Label          `json:"label"`
	LabelID        uint           `json:"-"`
	BeginTime      time.Time      `json:"beginTime"` // Refresh on start
	EndTime        time.Time      `json:"endTime"`   // Assign only after end
	Servant        User           `json:"servant"`   // Assign only after start
	ServantID      uint           `json:"-"`
	Meta           models.JSONMap `json:"meta"`
}
