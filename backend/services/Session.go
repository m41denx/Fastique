package services

type SessionService interface {
	NewSession(uid uint) (session string)
	DeleteSession(session string)
	GetSession(session string) (uid uint)
}
