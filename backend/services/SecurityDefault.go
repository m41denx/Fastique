package services

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
)

type DefaultSecurityService struct{}

func NewDefaultSecurityService() DefaultSecurityService {
	return DefaultSecurityService{}
}

func (DefaultSecurityService) ValidateUsername(username string) error {
	if len(username) < 4 || len(username) > 20 {
		return errors.New("Username must be between 4 and 20 characters")
	}
	return nil
}

func (DefaultSecurityService) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be 8 characters or more")
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return errors.New("Password must contain at least one lowercase character")
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return errors.New("Password must contain at least one uppercase character")
	}
	if !strings.ContainsAny(password, "0123456789") {
		return errors.New("Password must contain at least one digit")
	}
	return nil
}

func (DefaultSecurityService) HashPassword(password string) string {
	// sha256(sha1(password))
	s1 := sha1.New()
	s1.Write([]byte(password))
	s256 := sha256.New()
	s256.Write(s1.Sum(nil))
	return fmt.Sprintf("%x", s256.Sum(nil))
}
