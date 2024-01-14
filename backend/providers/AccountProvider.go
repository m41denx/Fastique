package providers

import (
	"gorm.io/gorm"
	"main/models/db"
	"main/services"
)

type AccountProvider struct {
	db       *gorm.DB
	sessions services.SessionService
}

func NewAccountProvider(db *gorm.DB, sessionStorage services.SessionService) *AccountProvider {
	return &AccountProvider{
		db:       db,
		sessions: sessionStorage,
	}
}

func (ap *AccountProvider) New() *Account {
	return &Account{
		p:    ap,
		User: &db.User{},
	}
}
