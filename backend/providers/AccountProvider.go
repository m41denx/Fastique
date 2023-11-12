package providers

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
	"main/models/db"
)

type AccountProvider struct {
	db       *gorm.DB
	sessions *session.Session
}

func NewAccountProvider(db *gorm.DB, sessionStorage *session.Session) *AccountProvider {
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
