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

func (t *TicketProvider) GetTicketsForDay(date time.Time, branchID ...uint) ([]db.Ticket, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	date2 := date.AddDate(0, 0, 1)
	var tickets []db.Ticket
	tx := t.db.Where("begin_time >= ? AND end_time <= ?", date, date2).
		Where("branch_id IN ?", branchID).
		Order("begin_time").Find(&tickets)
	return tickets, tx.Error
}

func (t *TicketProvider) GetTodayTickets(branchID ...uint) ([]db.Ticket, error) {
	return t.GetTicketsForDay(time.Now(), branchID...)
}

func (t *TicketProvider) GetAllTickets() ([]db.Ticket, error) {
	var tickets []db.Ticket
	tx := t.db.Model(&db.Ticket{}).Preload("Branch").Preload("Servant").Preload("Label").Find(&tickets)
	return tickets, tx.Error
}

func (t *TicketProvider) GetTodayTicketsForLabel(branchID uint, labelID uint) ([]db.Ticket, error) {
	tickets, err := t.GetTicketsForDay(time.Now(), branchID)
	if err != nil {
		return tickets, err
	}
	var tickets2 []db.Ticket
	for _, ticket := range tickets {
		if ticket.LabelID == labelID {
			tickets2 = append(tickets2, ticket)
		}
	}
	return tickets2, nil
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
