package providers

import (
	"gorm.io/gorm"
	"main/models/db"
	"time"
)

type TicketProvider struct {
	db *gorm.DB
}

func NewTicketProvider(db *gorm.DB) *TicketProvider {
	return &TicketProvider{
		db: db,
	}
}

func (t *TicketProvider) GetTicketsForDay(date time.Time) ([]db.Ticket, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	date2 := date.AddDate(0, 0, 1)
	var tickets []db.Ticket
	tx := t.db.Where("begin_time >= ?", date).Where("end_time <= ?", date2).Find(&tickets)
	return tickets, tx.Error
}

func (t *TicketProvider) GetTodayTickets() ([]db.Ticket, error) {
	return t.GetTicketsForDay(time.Now())
}

func (t *TicketProvider) GetTicket(id uint) (db.Ticket, error) {
	ticket := db.Ticket{}
	ticket.ID = id
	tx := t.db.First(&ticket)
	return ticket, tx.Error
}

func (t *TicketProvider) CreateTicket(ticket db.Ticket) (uint, error) {
	tx := t.db.Create(&ticket)
	return ticket.ID, tx.Error
}

func (t *TicketProvider) UpdateTicket(ticket db.Ticket) (uint, error) {
	tx := t.db.Updates(&ticket)
	return ticket.ID, tx.Error
}

func (t *TicketProvider) DeleteTicket(id uint) error {
	tx := t.db.Delete(&db.Ticket{}, id)
	return tx.Error
}
