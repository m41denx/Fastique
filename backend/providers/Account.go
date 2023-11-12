package providers

import (
	"main/models/db"
	"main/services"
)

type Account struct {
	p    *AccountProvider
	User *db.User
}

func (a *Account) Register(login string, password string) (uid uint, err error) {
	return a.RegisterWithSecService(login, password, services.NewDefaultSecurityService())
}

func (a *Account) RegisterWithSecService(login string, password string, sec services.SecurityService) (uid uint, err error) {
	err = sec.ValidateUsername(login)
	if err != nil {
		return 0, err
	}
	err = sec.ValidatePassword(password)
	if err != nil {
		return 0, err
	}

	a.User.Username = login
	a.User.PassHash = sec.HashPassword(password)
	tx := a.p.db.Create(a.User)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return a.User.ID, nil
}

func (a *Account) Login(login string, password string) (uid uint, err error) {
	return a.LoginWithSecService(login, password, services.NewDefaultSecurityService())
}

func (a *Account) LoginWithSecService(login string, password string, sec services.SecurityService) (uid uint, err error) {
	password = sec.HashPassword(password)
	err = a.p.db.Where(db.User{Username: login, PassHash: password}).First(&a.User).Error
	return a.User.ID, err
}

func (a *Account) GetUser(uid uint) (err error) {
	a.User.ID = uid
	return a.p.db.Preload("Role").First(&a.User).Error
}

func (a *Account) HasPermission(perm string) bool {
	return a.User.Role.Privileges[perm] == "1"
}
