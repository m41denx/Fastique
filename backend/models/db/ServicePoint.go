package db

import (
	"gorm.io/gorm"
)

type ServicePoint struct {
	gorm.Model
	BranchID  uint    `json:"-"`
	Labels    []Label `json:"labels" gorm:"many2many:sp_labels"`
	Available bool    `json:"available"`
	User      User    `json:"user"`
	UserID    uint    `json:"-"`
}
