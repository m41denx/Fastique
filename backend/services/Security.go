package services

type SecurityService interface {
	ValidateUsername(username string) error
	ValidatePassword(password string) error
	HashPassword(password string) string
}
