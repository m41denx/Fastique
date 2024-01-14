package providers

import (
	"gorm.io/gorm"
	"main/models/db"
)

type ServicePointProvider struct {
	db *gorm.DB
	sp *db.ServicePoint
}

func NewServicePointProvider(db *gorm.DB) *ServicePointProvider {
	return &ServicePointProvider{
		db: db,
	}
}

func (s *ServicePointProvider) LoadServicePoint(id uint) error {
	s.sp.ID = id
	return s.db.First(&s.sp).Error
}

func (s *ServicePointProvider) AuthUser(user db.User) bool {
	if s.sp.User.ID != 0 {
		return false
	}
	s.sp.User = user
	return s.db.Updates(&s.sp).Error == nil
}

func (s *ServicePointProvider) DeauthUser(user db.User) bool {
	if s.sp.User.ID != user.ID {
		return false
	}
	s.sp.User = db.User{}
	return s.db.Updates(&s.sp).Error == nil
}

func (s *ServicePointProvider) AddLabel(label db.Label) error {
	s.sp.Labels = append(s.sp.Labels, label)
	return s.db.Updates(&s.sp).Error
}

func (s *ServicePointProvider) RemoveLabel(label db.Label) error {
	for i, l := range s.sp.Labels {
		if l.ID == label.ID {
			s.sp.Labels = append(s.sp.Labels[:i], s.sp.Labels[i+1:]...)
			return s.db.Updates(&s.sp).Error
		}
	}
	return nil
}

func (s *ServicePointProvider) SetAvailability(available bool) error {
	s.sp.Available = available
	return s.db.Updates(&s.sp).Error
}

func (s *ServicePointProvider) GetTickets() []db.Ticket {
	var tickets []db.Ticket
	s.db.Where(db.Ticket{ServicePointID: s.sp.ID}).Find(&tickets)
	return tickets
}

//func (s *ServicePointProvider) ClaimTicket(ticket db.Ticket) bool {
//	s.sp.Tickets = append(s.sp.Tickets, ticket)
//	return s.db.Updates(&s.sp).Error == nil
//}
//
//func (s *ServicePointProvider) ReleaseTicket(ticket db.Ticket) {
//	for i, t := range s.sp.Tickets {
//		if t.ID == ticket.ID {
//			s.sp.Tickets = append(s.sp.Tickets[:i], s.sp.Tickets[i+1:]...)
//			return
//		}
//	}
//}
