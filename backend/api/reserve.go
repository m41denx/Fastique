package api

import (
	"github.com/gofiber/fiber/v2"
	"main/models"
	"main/models/db"
	"strconv"
	"strings"
	"time"
)

func (api *API) Reserve(c *fiber.Ctx) error {
	d := &struct {
		LabelID  uint      `json:"label"`
		BranchID uint      `json:"branch"`
		Date     time.Time `json:"date"`
	}{}

	if err := c.BodyParser(d); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	label := db.Label{}
	if err := api.OrgProvider.ExposeGorm().First(&label, d.LabelID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	branch := db.Branch{}
	if err := api.OrgProvider.ExposeGorm().First(&branch, d.BranchID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	daydate := time.Date(d.Date.Year(), d.Date.Month(), d.Date.Day(), 0, 0, 0, 0, time.UTC)

	var tickets int64

	if err := api.OrgProvider.ExposeGorm().Model(&db.Ticket{}).
		Where("begin_time >= ? AND end_time <= ? AND label_id = ?", daydate, daydate.AddDate(0, 0, 1), d.LabelID).
		Count(&tickets).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}

	ticket := db.Ticket{
		Name:      strings.ReplaceAll(label.Template, "#", strconv.Itoa(int(tickets+1))),
		BranchID:  d.BranchID,
		LabelID:   d.LabelID,
		BeginTime: d.Date,
	}

	if _, err := api.TicketProvider.CreateTicket(ticket); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiError{Message: err.Error()})
	}
	return c.JSON(fiber.Map{"order": tickets + 1, "load": tickets})
}
