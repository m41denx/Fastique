package services

import "github.com/google/uuid"

type SessionBasicService struct {
	sessions map[string]uint
}

func NewSessionBasicService() *SessionBasicService {
	return &SessionBasicService{
		sessions: make(map[string]uint),
	}
}

func (s *SessionBasicService) NewSession(uid uint) (session string) {
	session = uuid.NewString()
	s.sessions[session] = uid
	return session
}

func (s *SessionBasicService) DeleteSession(session string) {
	delete(s.sessions, session)
}

func (s *SessionBasicService) GetSession(session string) (uid uint) {
	return s.sessions[session]
}
